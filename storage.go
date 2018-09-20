package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Db interface {
	GetPosts() Posts
	GetByTerm(searchTerm string) Posts
	GetById(PostId string) Post
	User() string
}

type DB struct {
	user,
	password,
	name string
}

func (m *DB) User() string {
	return m.user
}

func (m *DB) dbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", m.user+":"+m.password+"@/"+m.name+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func (m *DB) GetPosts() Posts {
	db := m.dbConn()
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
	return posts
}

func (m *DB) GetByTerm(searchTerm string) Posts {
	post := Post{}
	posts := Posts{}

	db := m.dbConn()
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
	return posts
}

func (m *DB) GetById(postId string) Post {
	post := Post{}

	db := m.dbConn()
	defer db.Close()
	row, err := db.Query("SELECT * FROM posts WHERE id=?", postId)
	if err != nil {
		panic(err.Error())
	}

	result := row.Next()
	if !result {
		return post
	}

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

	return post
}
