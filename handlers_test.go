package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestBlog(t *testing.T) {

	db := dbConn()
	defer db.Close()

	if _, err := db.Exec("TRUNCATE TABLE posts"); err != nil {
		t.Fatal(err)
	}

	res, err := db.Exec("INSERT INTO posts (title, body) VALUES ('post title', 'post body')")
	if err != nil {
		t.Fatal(err)
	}
	id, _ := res.LastInsertId()

	router := NewRouter()
	ts := httptest.NewServer(router)
	defer ts.Close()

	var cases = []struct {
		handlerFunc    http.HandlerFunc
		path           string
		expectedStatus int
	}{
		{
			Index,
			ts.URL + "/",
			http.StatusOK,
		},
		{
			Show,
			fmt.Sprintf("%s/post/%d", ts.URL, id),
			http.StatusOK,
		},
		{
			Show,
			fmt.Sprintf("%s/post/%d", ts.URL, rand.Int()),
			http.StatusNotFound,
		},
		{
			Search,
			ts.URL + "/search?q='title'",
			http.StatusOK,
		},
	}

	for _, tt := range cases {
		resp, err := http.Get(tt.path)
		if err != nil {
			t.Fatal(err)
		}

		status := resp.StatusCode

		if status != tt.expectedStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, tt.expectedStatus)
		}
	}
}

func TestSearch(t *testing.T) {
	title := strconv.Itoa(rand.Int())

	db := dbConn()
	defer db.Close()

	if _, err := db.Exec("TRUNCATE TABLE posts"); err != nil {
		t.Fatal(err)
	}

	_, err := db.Exec("INSERT INTO posts (title, body) VALUES (?, 'random post body')", title)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	var cases = []struct {
		postTitle     string
		expectedValue bool
	}{
		{
			title,
			true,
		},
		{
			strconv.Itoa(rand.Int()),
			false,
		},
	}

	for _, tt := range cases {
		req, err := http.NewRequest("GET", "/search?q="+tt.postTitle, nil)
		if err != nil {
			t.Fatal(err)
		}

		handler := http.HandlerFunc(Search)
		handler.ServeHTTP(rr, req)
		if strings.Contains(rr.Body.String(), tt.postTitle) != tt.expectedValue {
			t.Errorf("Uncorrect titles returned")
		}
	}
}
