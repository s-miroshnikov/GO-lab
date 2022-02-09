package main

import (
	
	"database/sql"
	
	"log"
	_ "github.com/jackc/pgx/v4/stdlib"
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
	type Articles struct {
		id        int64
		magazines_id	int64
		article_type_id	int64
		author_id	int64
	  }
	  
	  type magazines struct {
		id   int64
		name string
	  }
	  
	  type article_types struct {
		id      int64
		"type"    string
	  }
	  type author struct {
		id   int64
		author string
	  }
}
