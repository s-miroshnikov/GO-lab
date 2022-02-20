package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/alexeyco/pig"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type Magazine struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

type Article struct {
	ID              int64 `db:"id"`
	Magazine_ID     int64 `db:"magazines_id"`
	Article_Type_ID int64 `db:"article_type_id"`
	Author_ID       int64 `db:"author_id"`
}

type Article_Type struct {
	ID   int64  `db:"id"`
	Type string `db:"type1"`
}

type Author struct {
	ID   int64  `db:"id"`
	Type string `db:"author"`
}

func main() {

	conn, err := pgx.Connect(context.Background(), os.Getenv("PGDSN"))
	if err != nil {
		log.Fatalln(err)
	}

	p := pig.New(conn)

	var magazines []Magazine
	err = p.Query().Select(&magazines, "SELECT * FROM magazines")
	if err != nil {
		log.Fatalln(err)
	}
	var author []Author
	err = p.Query().Select(&author, "SELECT * FROM author")
	if err != nil {
		log.Fatalln(err)
	}
	var article_types []Article_Type
	err = p.Query().Select(&article_types, "SELECT * FROM article_types")
	if err != nil {
		log.Fatalln(err)
	}
	var articles []Article
	err = p.Query().Select(&articles, "SELECT * FROM articles")
	if err != nil {
		log.Fatalln(err)
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.LoadHTMLGlob("/local/scripts/templates/*.tmpl")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "library.tmpl", gin.H{
			"title":         "DB library",
			"Articles":      articles,
			"Article_types": article_types,
			"Magazines":     magazines,
			"Authors":       author,
		})
	})
	router.Run(":8080")
}
