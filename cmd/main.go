package main

import (
	"fmt"

	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/repository"
)

func main() {
	db, err := repository.InitializeDb()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Database connection successful!")
	}
	defer db.Close()
}
