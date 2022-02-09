package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db, err := sql.Open("pgx", os.Getenv("PGDSN"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	type Articles struct {
		id              int64
		magazines_id    int64
		article_type_id int64
		author_id       int64
	}

	type magazines struct {
		id   int64
		name string
	}

	type article_types struct {
		id    int64
		type1 string
	}
	type author struct {
		id     int64
		author string
	}
}
