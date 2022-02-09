package main

import (
	
	"database/sql"
	
	"log"
	"github.com/jackc/pgx/v4"
)

func main() {
	connectionString := "host=db user=worker1 password=xxxxx dbname=library sslmode=disable"
	db, err := sql.Open( driverName: "pgx", connectionString)
    if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}
