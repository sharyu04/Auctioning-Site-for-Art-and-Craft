package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/app"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/middleware"
)

func NewRouter(deps app.Dependencies) mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server started"))
	})

	router.HandleFunc("/user/signup", createUserHandler(deps.UserService)).Methods("POST")
	router.HandleFunc("/admin/signup", createAdminHandler(deps.UserService)).Methods("POST")
	router.HandleFunc("/login", loginHandler(deps.UserService)).Methods("POST")

	router.Handle("/check", middleware.RequireAuth(checkHandler))

	router.HandleFunc("/artwork/create", createArtworkHandler(deps.ArtworkService)).Methods("POST")
	router.HandleFunc("/artworks", GetArtworksHandler(deps.ArtworkService)).Methods("GET")
	router.HandleFunc("/artwork/{id}", GetArtworkByIdHandler(deps.ArtworkService)).Methods("GET")

	router.HandleFunc("/bid/create", createBidHandler(deps.BidService)).Methods("POST")

	return *router
}

var checkHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	user_id := r.Header.Get("user_id")
	role := r.Header.Get("role")
	fmt.Printf("\n User Id: %s", user_id)
	fmt.Printf("\n Role: %s", role)
	w.Write([]byte("Authorization Successful!"))
})
