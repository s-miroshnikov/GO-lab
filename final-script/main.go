package main

import (
	"context"
	"io"
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
	fn := logOutput()
	defer fn()
	conn, err := pgx.Connect(context.Background(), "postgres://worker1:Q23fs12qwe@db/library")
	if err != nil {
		log.Fatalln(err)
	}
	p := pig.New(conn)

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.LoadHTMLGlob("/local/scripts/templates/*.tmpl")

	router.GET("/", func(c *gin.Context) {
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

func logOutput() func() {
	logfile := `logfile`
	// open file read/write | create if not exist | clear file at open if exists
	f, _ := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)

	// save existing stdout | MultiWriter writes to saved stdout and file
	out := os.Stdout
	mw := io.MultiWriter(out, f)

	// get pipe reader and writer | writes to pipe writer come out pipe reader
	r, w, _ := os.Pipe()

	// replace stdout,stderr with pipe writer | all writes to stdout, stderr will go through pipe instead (fmt.print, log)
	os.Stdout = w
	os.Stderr = w

	// writes with log.Print should also write to mw
	log.SetOutput(mw)

	//create channel to control exit | will block until all copies are finished
	exit := make(chan bool)

	go func() {
		// copy all reads from pipe to multiwriter, which writes to stdout and file
		_, _ = io.Copy(mw, r)
		// when r or w is closed copy will finish and true will be sent to channel
		exit <- true
	}()

	// function to be deferred in main until program exits
	return func() {
		// close writer then block on exit channel | this will let mw finish writing before the program exits
		_ = w.Close()
		<-exit
		// close file after all writes have finished
		_ = f.Close()
	}

}
