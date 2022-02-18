package main

import (
	"context"
	"log"
	"os"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Articles struct {
	id              int64
	magazines_id    int64
	article_type_id int64
	author_id       int64
}

type Magazines struct {
	id   int64
	name string
}

type Article_Types struct {
	id    int64
	type1 string
}
type Author struct {
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
	conn, _ := pgxpool.Connect(ctx, os.Getenv("PGDSN"))

	var articles []*Articles
	pgxscan.Select(ctx, conn, &articles, `SELECT id, magazines_id, article_type_id, author_id FROM Articles`)
	var magazines []*Magazines
	pgxscan.Select(ctx, conn, &magazines, `SELECT id, name FROM magazines`)
	var article_types []*Article_Types
	pgxscan.Select(ctx, conn, &article_types, `SELECT id, type FROM article_types`)
	var author []*Author
	pgxscan.Select(ctx, conn, &author, `SELECT id, author FROM author`)
}
