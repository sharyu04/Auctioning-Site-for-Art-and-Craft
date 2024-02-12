package api

import (
	// "encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/app"

	// "github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/app/artwork"
	// "github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/dto"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/middleware"
)

func NewRouter(deps app.Dependencies) mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server started"))
	})

	router.HandleFunc("/user/signup", createUserHandler(deps.UserService)).Methods("POST")
	router.Handle("/admin/signup", middleware.RequireAuth(createAdminHandler(deps.UserService), []string{"super_admin", "admin"})).Methods("POST")
	router.HandleFunc("/login", loginHandler(deps.UserService)).Methods("POST")

	router.Handle("/check", middleware.RequireAuth(checkHandler, []string{"user", "admin", "super_admin"}))

	router.Handle("/artwork/create", middleware.RequireAuth(createArtworkHandler(deps.ArtworkService), []string{"admin", "user"})).Methods("POST")
	router.HandleFunc("/artworks", GetArtworksHandler(deps.ArtworkService)).Methods("GET")        //middleware
	router.HandleFunc("/artwork/{id}", GetArtworkByIdHandler(deps.ArtworkService)).Methods("GET") //middleware
	router.Handle("/artwork/{id}", middleware.RequireAuth(DeleteArtworkHandler(deps.ArtworkService), []string{"super_admin", "admin", "user"})).Methods("DELETE")

	router.Handle("/bid/create", middleware.RequireAuth(createBidHandler(deps.BidService), []string{"user"})).Methods("POST")
	router.Handle("/bid/update", middleware.RequireAuth(updateBidHandler(deps.BidService), []string{"user"})).Methods("PATCH") //PUT

	return *router
}

var checkHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	user_id := r.Header.Get("user_id")
	role := r.Header.Get("role")
	fmt.Printf("\n User Id: %s", user_id)
	fmt.Printf("\n Role: %s", role)
	w.Write([]byte("Authorization Successful!"))
})
