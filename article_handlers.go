package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := dbGetAllArticles()
	catch(err)

	t, _ := template.ParseFiles("templates/base.html", "templates/index.html")
	err = t.Execute(w, articles)
	catch(err)
}

func NewArticle(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/base.html", "templates/new.html")
	err := t.Execute(w, nil)
	catch(err)
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	content := r.FormValue("content")
	article := &Article{Title: title, Content: template.HTML(content)}

	err := dbCreateArticle(article)
	catch(err)
	http.Redirect(w, r, "/", http.StatusFound)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	article := r.Context().Value("article").(*Article)
	t, _ := template.ParseFiles("templates/base.html", "templates/article.html")
	err := t.Execute(w, article)
	catch(err)
}

func EditArticle(w http.ResponseWriter, r *http.Request) {
	article := r.Context().Value("article").(*Article)

	t, _ := template.ParseFiles("templates/base.html", "templates/edit.html")
	err := t.Execute(w, article)
	catch(err)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	article := r.Context().Value("article").(*Article)

	title := r.FormValue("title")
	content := r.FormValue("content")
	newArticle := &Article{Title: title, Content: template.HTML(content)}

	err := dbUpdateArticle(strconv.Itoa(article.ID), newArticle)
	catch(err)
	http.Redirect(w, r, fmt.Sprintf("/articles/%d", article.ID), http.StatusFound)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	article := r.Context().Value("article").(*Article)
	err := dbDeleteArticle(strconv.Itoa(article.ID))
	catch(err)

	http.Redirect(w, r, "/", http.StatusFound)
}
