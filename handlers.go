package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var tmpl = template.Must(template.ParseGlob("templates/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM posts ORDER BY created DESC")
	if err != nil {
		panic(err.Error())
	}

	post := Post{}
	posts := Posts{}
	for rows.Next() {
		var id int
		var title, body string
		var created time.Time
		err = rows.Scan(&id, &title, &body, &created)
		if err != nil {
			panic(err.Error())
		}
		post.Id = id
		post.Title = title
		post.Body = body
		post.Date = post.FormatDate(created)
		posts = append(posts, post)
	}
	err = tmpl.ExecuteTemplate(w, "index.gohtml", posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Search(w http.ResponseWriter, r *http.Request) {
	post := Post{}
	posts := Posts{}

	terms, ok := r.URL.Query()["q"]
	if !ok || len(terms[0]) < 1 {
		json.NewEncoder(w).Encode(posts)
		return
	}
	searchTerm := terms[0]
	db := dbConn()
	defer db.Close()

	query := `SELECT id, title FROM posts WHERE posts.title LIKE CONCAT('%', ?, '%')`
	rows, err := db.Query(query, searchTerm)
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var id int
		var title string
		err = rows.Scan(&id, &title)
		if err != nil {
			panic(err.Error())
		}
		post.Id = id
		post.Title = title
		posts = append(posts, post)
	}
	json.NewEncoder(w).Encode(posts)
}

func Show(w http.ResponseWriter, r *http.Request) {
	postId := mux.Vars(r)["id"]

	db := dbConn()
	defer db.Close()
	row, err := db.Query("SELECT * FROM posts WHERE id=?", postId)
	if err != nil {
		panic(err.Error())
	}

	result := row.Next()
	if !result {
		http.Error(w, "The post you are looking for doesn't exist", http.StatusNotFound)
		return
	}

	post := Post{}
	var id int
	var title, body string
	var created time.Time
	err = row.Scan(&id, &title, &body, &created)
	if err != nil {
		panic(err.Error())
	}
	post.Id = id
	post.Title = title
	post.Body = body
	post.Date = post.FormatDate(created)

	err = tmpl.ExecuteTemplate(w, "show.gohtml", post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
