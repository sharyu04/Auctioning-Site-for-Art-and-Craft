package middleware

import (
	"net/http"
	"slices"

	"github.com/golang-jwt/jwt"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/dto"
)

func RequireAuth(next http.Handler, roles []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var jwtKey = []byte("secret_key")

		cookie, err := r.Cookie("accessToken")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized, Auth Token not found!"))
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		tokenString := cookie.Value

		claims := &dto.Claims{}

		tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, err
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}

		Id := claims.Id
		Role := claims.Role

		if !slices.Contains(roles, Role) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Role unauthorized"))
			return
		}

		r.Header.Set("user_id", Id.String())
		r.Header.Set("role", Role)

		next.ServeHTTP(w, r)

	})

}
