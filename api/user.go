package api

import (
	"encoding/json"
	"net/http"

	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/app/user"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/dto"
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
