package main

import (
	"fmt"
	"net/http"

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
	http.ListenAndServe("localhost:8080", &router)
}
