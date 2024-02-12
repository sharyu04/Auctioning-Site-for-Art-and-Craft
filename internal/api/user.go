package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/app/user"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/dto"
)

func createUserHandler(userSvc user.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.CreateUserRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		resBody, err := userSvc.CreateUser(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		respJson, err := json.Marshal(resBody)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(respJson)
		return
	}
}

func createAdminHandler(userSvc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role := r.Header.Get("role")
		fmt.Println(role)
		if role != "super_admin" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized to register admin"))
			return
		}

		var req dto.CreateUserRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		resBody, err := userSvc.CreateAdmin(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		respJson, err := json.Marshal(resBody)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(respJson)
		return
	}
}

func loginHandler(userSvc user.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.LoginRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		if len(req.Email) == 0 || len(req.Password) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Please provide name and password to login"))
			return
		}

		token, err := userSvc.LoginUser(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		cookie := http.Cookie{}
		cookie.Name = "accessToken"
		cookie.Value = token
		cookie.Expires = time.Now().Add(time.Minute * 5)
		cookie.Secure = false
		cookie.HttpOnly = true
		cookie.Path = "/"
		http.SetCookie(w, &cookie)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(token))
		return
	}
}
