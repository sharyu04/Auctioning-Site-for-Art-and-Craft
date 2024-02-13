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

	//user routes
	router.HandleFunc("/user/signup", createUserHandler(deps.UserService)).Methods("POST")
	router.Handle("/admin/signup", middleware.RequireAuth(createUserHandler(deps.UserService), []string{"super_admin", "admin"})).Methods("POST")
	router.HandleFunc("/login", loginHandler(deps.UserService)).Methods("POST")
	router.Handle("/users", middleware.RequireAuth(GetAllUsersHandler(deps.UserService), []string{"admin", "super_admin"})).Methods("Get")

	router.Handle("/check", middleware.RequireAuth(checkHandler, []string{"user", "admin", "super_admin"}))

	//artwork routes
	router.Handle("/artwork/create", middleware.RequireAuth(createArtworkHandler(deps.ArtworkService), []string{"admin", "user"})).Methods("POST")
	router.Handle("/artworks", middleware.RequireAuth(GetArtworksHandler(deps.ArtworkService), []string{"super_admin", "admin", "user"})).Methods("GET")
	router.Handle("/artwork/{id}", middleware.RequireAuth(GetArtworkByIdHandler(deps.ArtworkService), []string{"super_admin", "admin", "user"})).Methods("GET")
	router.Handle("/artwork/{id}", middleware.RequireAuth(DeleteArtworkHandler(deps.ArtworkService), []string{"super_admin", "admin", "user"})).Methods("DELETE")

	//bids routes
	router.Handle("/bid/create", middleware.RequireAuth(createBidHandler(deps.BidService), []string{"user"})).Methods("POST")
	router.Handle("/bid/update", middleware.RequireAuth(updateBidHandler(deps.BidService), []string{"user"})).Methods("PUT")

	return *router
}

var checkHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	user_id := r.Header.Get("user_id")
	role := r.Header.Get("role")
	fmt.Printf("\n User Id: %s", user_id)
	fmt.Printf("\n Role: %s", role)
	w.Write([]byte("Authorization Successful!"))
})
