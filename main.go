package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type Blog struct {
	router *mux.Router
	srv    *http.Server
	db     Db
}

func NewRouter(routes Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	return router
}

func NewBlog(dbUser, dbPass, dbName string) *Blog {
	b := &Blog{}
	b.router = NewRouter(b.GetRoutes())
	b.srv = &http.Server{
		Handler:      b.router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	b.db = &DB{
		user:     dbUser,
		password: dbPass,
		name:     dbName,
	}
	return b
}

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	b := NewBlog(dbUser, dbPass, dbName)
	fmt.Println(b)
	log.Fatal(b.srv.ListenAndServe())
}
