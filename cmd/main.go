package main

import (
	"fmt"
	"net/http"
	"github.com/rs/cors"
	// "github.com/gorilla/handlers"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/api"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/app"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/repository"
)

func main() {
	db, err := repository.InitializeDb()
	// Port := 8080
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Database connection successful!")
	}
	defer db.Close()

	sevices := app.NewServices(db)

	router := api.NewRouter(sevices)

	// cors := handlers.CORS(
    //     handlers.AllowedHeaders([]string{"*"}),
    //     handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
    //     handlers.AllowedOrigins([]string{"*"}),
    // )
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"*"},
	})

	handler := c.Handler(&router)
	http.ListenAndServe(":8080", handler)
	
    // http.ListenAndServe(":8080", cors(&router))
	
}
