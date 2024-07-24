package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

var router *chi.Mux
var db *sql.DB

func main() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)

	var err error
	db, err = connect()
	catch(err)

	router.Use(ChangeMethod)
	router.Get("/", GetAllArticles)
	router.Post("/upload", UploadHandler)
	router.Get("/images/*", ServeImages)
	router.Route("/articles", func(r chi.Router) {
		r.Get("/", NewArticle)
		r.Post("/", CreateArticle)
		r.Route("/{articleID}", func(r chi.Router) {
			r.Use(ArticleCtx)
			r.Get("/", GetArticle)
			r.Put("/", UpdateArticle)
			r.Delete("/", DeleteArticle)
			r.Get("/edit", EditArticle)
		})
	})

	fmt.Println("Listening on port 8005")
	err = http.ListenAndServe(":8005", router)
	catch(err)
}

func catch(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
