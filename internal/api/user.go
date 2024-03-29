package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/app/user"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/apperrors"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/dto"
)

func createUserHandler(userSvc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req dto.CreateUserRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			errResponse := apperrors.MapError(err)
			w.WriteHeader(errResponse.ErrorCode)
			res, _ := json.Marshal(errResponse)
			w.Write(res)
			return
		}

		role := r.Header.Get("role")

		resBody, err := userSvc.CreateUser(req, role)
		if err != nil {
			errResponse := apperrors.MapError(err)
			w.WriteHeader(errResponse.ErrorCode)
			res, _ := json.Marshal(errResponse)
			w.Write(res)
			return
		}

		respJson, err := json.Marshal(resBody)
		if err != nil {
			errResponse := apperrors.MapError(err)
			w.WriteHeader(errResponse.ErrorCode)
			res, _ := json.Marshal(errResponse)
			w.Write(res)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(respJson)
	}
}

type resBody struct {
	token  string
	userId uuid.UUID
	role   string
}

func loginHandler(userSvc user.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.LoginRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			errResponse := apperrors.MapError(err)
			w.WriteHeader(errResponse.ErrorCode)
			res, _ := json.Marshal(errResponse)
			w.Write(res)
			return
		}

		if len(req.Email) == 0 || len(req.Password) == 0 {
			err := apperrors.BadRequest{ErrorMsg: "Please provide name and password to login"}
			errResponse := apperrors.MapError(err)
			w.WriteHeader(errResponse.ErrorCode)
			res, _ := json.Marshal(errResponse)
			w.Write(res)
			return
		}

		token, usesrId, userRole, err := userSvc.LoginUser(req)
		if err != nil {
			errResponse := apperrors.MapError(err)
			w.WriteHeader(errResponse.ErrorCode)
			res, _ := json.Marshal(errResponse)
			w.Write(res)
			return
		}

		cookie := http.Cookie{
			Name:     "accessToken",
			Value:    token,
			SameSite: http.SameSiteLaxMode,
			Path:     "/",
			MaxAge:   1000,
			HttpOnly: true,
			Secure:   false,
		}

		http.SetCookie(w, &cookie)

		// resBody := map[string]string{
		// 	"auth-token": token,
		// }

		resBody := resBody{
			token:  token,
			userId: usesrId,
			role:   userRole,
		}

		res, _ := json.Marshal(resBody)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func logoutHandler(userSvc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		resBody := map[string]string{
			"Response": "Logout successful!",
		}
		res, _ := json.Marshal(resBody)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func GetAllUsersHandler(userSvc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := r.URL.Query().Get("start")
		count := r.URL.Query().Get("count")
		role := r.URL.Query().Get("role")

		if start == "" || count == "" {
			err := apperrors.BadRequest{ErrorMsg: "start and count values missing!"}
			errResponse := apperrors.MapError(err)
			w.WriteHeader(errResponse.ErrorCode)
			res, _ := json.Marshal(errResponse)
			w.Write(res)
			return
		}

		startInt, _ := strconv.Atoi(start)
		countInt, _ := strconv.Atoi(count)

		res, err := userSvc.GetAllUsers(startInt, countInt, role)
		if err != nil {
			errResponse := apperrors.MapError(err)
			w.WriteHeader(errResponse.ErrorCode)
			res, _ := json.Marshal(errResponse)
			w.Write(res)
			return
		}

		resBody, err := json.Marshal(res)
		if err != nil {
			errResponse := apperrors.MapError(err)
			w.WriteHeader(errResponse.ErrorCode)
			res, _ := json.Marshal(errResponse)
			w.Write(res)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(resBody)

	}
}
