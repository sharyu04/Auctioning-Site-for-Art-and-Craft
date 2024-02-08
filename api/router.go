package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/app"
)

func NewRouter(deps app.Dependencies) mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server started"))
	})

	router.HandleFunc("/user/signup", createUserHandler(deps.UserService)).Methods("POST")

	return *router
}
