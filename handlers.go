package main

import (
	"encoding/json"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var tmpl = template.Must(template.ParseGlob("templates/*"))

func (b *Blog) Index(w http.ResponseWriter, r *http.Request) {
	posts := b.db.GetPosts()
	if err := tmpl.ExecuteTemplate(w, "index.gohtml", posts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (b *Blog) Search(w http.ResponseWriter, r *http.Request) {
	posts := Posts{}
	terms, ok := r.URL.Query()["q"]
	if !ok || len(terms[0]) < 1 {
		json.NewEncoder(w).Encode(posts)
		return
	}
	searchTerm := terms[0]
	posts = b.db.GetByTerm(searchTerm)
	json.NewEncoder(w).Encode(posts)
	return
}

func (b *Blog) Show(w http.ResponseWriter, r *http.Request) {
	postId := mux.Vars(r)["id"]
	post := b.db.GetById(postId)
	if post.IsEmpty() {
		http.Error(w, "The post you are looking for doesn't exist", http.StatusNotFound)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "show.gohtml", post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}
