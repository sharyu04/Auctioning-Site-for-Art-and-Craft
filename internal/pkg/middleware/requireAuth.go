package middleware

import (
	"encoding/json"
	"net/http"
	"slices"
	
	"github.com/golang-jwt/jwt"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/apperrors"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/dto"
)

func RequireAuth(next http.Handler, roles []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var jwtKey = []byte("secret_key")

		cookie, err := r.Cookie("accessToken")
		if err != nil {
			if err == http.ErrNoCookie {
				err = apperrors.UnAuthorizedAccess{ErrorMsg: "Unauthorized, Auth Token not found!"}
							errResponse := apperrors.MapError(err)
				w.WriteHeader(errResponse.ErrorCode)
				res, _ := json.Marshal(errResponse)
				w.Write(res)
				return
			}
			err = apperrors.UnAuthorizedAccess{ErrorMsg: err.Error()}
			errResponse := apperrors.MapError(err)
			w.WriteHeader(errResponse.ErrorCode)
			res, _ := json.Marshal(errResponse)
			w.Write(res)
			return
		}

		tokenString := cookie.Value

		claims := &dto.Claims{}

		tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, err
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				err = apperrors.UnAuthorizedAccess{ErrorMsg: err.Error()}
				errResponse := apperrors.MapError(err)
				w.WriteHeader(errResponse.ErrorCode)
				res, _ := json.Marshal(errResponse)
				w.Write(res)
				return
			}
			err = apperrors.UnAuthorizedAccess{ErrorMsg: err.Error()}
			errResponse := apperrors.MapError(err)
			w.WriteHeader(errResponse.ErrorCode)
			res, _ := json.Marshal(errResponse)
			w.Write(res)
			return
		}

		if !tkn.Valid {
			err := apperrors.UnAuthorizedAccess{ErrorMsg: "Token Invalid"}
			errResponse := apperrors.MapError(err)
			w.WriteHeader(errResponse.ErrorCode)
			res, _ := json.Marshal(errResponse)
			w.Write(res)
			return
		}

		Id := claims.Id
		Role := claims.Role

		if !slices.Contains(roles, Role) {
			err := apperrors.UnAuthorizedAccess{ErrorMsg: "Role Unauthorized"}
			errResponse := apperrors.MapError(err)
			w.WriteHeader(errResponse.ErrorCode)
			res, _ := json.Marshal(errResponse)
			w.Write(res)
			return
		}

		r.Header.Set("user_id", Id.String())
		r.Header.Set("role", Role)

		next.ServeHTTP(w, r)

	})

}
