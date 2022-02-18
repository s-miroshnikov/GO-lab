package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

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

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, os.Getenv("PGDSN"))
	Check(err)
	defer conn.Close(ctx)

}

func GetAllArticles(conn *pgx.Conn) {
	// Execute the query
	if rows, err := conn.Query(context.Background(), "SELECT * FROM Articles"); err != nil {
		fmt.Println("Unable to query due to: ", err)
		return
	} else {
		// carefully deferring Queries closing
		defer rows.Close()

		// Using tmp variable for reading
		var tmp Articles

		// Next prepares the next row for reading.
		for rows.Next() {
			// Scan reads the values from the current row into tmp
			rows.Scan(&tmp.id, &tmp.magazines_id, &tmp.article_type_id, &tmp.author_id)
			fmt.Printf("%+v\n", tmp)
		}
		if rows.Err() != nil {
			// if any error occurred while reading rows.
			fmt.Println("Error will reading Articles table: ", err)
			return
		}
	}
}

func GetAllMagazines(conn *pgx.Conn) {
	// Execute the query
	if rows, err := conn.Query(context.Background(), "SELECT * FROM magazines"); err != nil {
		fmt.Println("Unable to query due to: ", err)
		return
	} else {
		// carefully deferring Queries closing
		defer rows.Close()

		// Using tmp variable for reading
		var tmp magazines

		// Next prepares the next row for reading.
		for rows.Next() {
			// Scan reads the values from the current row into tmp
			rows.Scan(&tmp.id, &tmp.name)
			fmt.Printf("%+v\n", tmp)
		}
		if rows.Err() != nil {
			// if any error occurred while reading rows.
			fmt.Println("Error will reading magazines table: ", err)
			return
		}
	}
}

func GetAllArticleTypes(conn *pgx.Conn) {
	// Execute the query
	if rows, err := conn.Query(context.Background(), "SELECT * FROM article_types"); err != nil {
		fmt.Println("Unable to query due to: ", err)
		return
	} else {
		// carefully deferring Queries closing
		defer rows.Close()

		// Using tmp variable for reading
		var tmp article_types

		// Next prepares the next row for reading.
		for rows.Next() {
			// Scan reads the values from the current row into tmp
			rows.Scan(&tmp.id, &tmp.type1)
			fmt.Printf("%+v\n", tmp)
		}
		if rows.Err() != nil {
			// if any error occurred while reading rows.
			fmt.Println("Error will reading article_types table: ", err)
			return
		}
	}
}

func GetAllAuthors(conn *pgx.Conn) {
	// Execute the query
	if rows, err := conn.Query(context.Background(), "SELECT * FROM author"); err != nil {
		fmt.Println("Unable to query due to: ", err)
		return
	} else {
		// carefully deferring Queries closing
		defer rows.Close()

		// Using tmp variable for reading
		var tmp author

		// Next prepares the next row for reading.
		for rows.Next() {
			// Scan reads the values from the current row into tmp
			rows.Scan(&tmp.id, &tmp.author)
			fmt.Printf("%+v\n", tmp)
		}
		if rows.Err() != nil {
			// if any error occurred while reading rows.
			fmt.Println("Error will reading author table: ", err)
			return
		}
	}
}
