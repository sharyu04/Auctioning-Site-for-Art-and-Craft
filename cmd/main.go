package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
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

	cors := handlers.CORS(
        handlers.AllowedHeaders([]string{"*"}),
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
        handlers.AllowedOrigins([]string{"*"}),
    )

    http.ListenAndServe(":8080", cors(&router))
	
}
