package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/app"
)

func NewRouter(deps app.Dependencies) mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server started"))
	})

	router.HandleFunc("/user/signup", createUserHandler(deps.UserService)).Methods("POST")
	// router.HandleFunc("/login", loginHandler(deps.UserService)).Methods("POST")

	router.HandleFunc("/artwork/create", createArtworkHandler(deps.ArtworkService)).Methods("POST")

	return *router
}
