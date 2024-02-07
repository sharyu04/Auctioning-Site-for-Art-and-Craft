package repository

import (
	// "database/sql"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "AuctionWebsite"
)

func InitializeDb() {
	psqlconn := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode = disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Database connection successful!")
	}
	defer db.Close()
}
