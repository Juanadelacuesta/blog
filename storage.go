package storage

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Db interface {
	getPosts() 
}

var Mysql struct {
	user,
	password,
	name string
}

db := dbConn()
defer db.Close()

rows, err := db.Query("SELECT * FROM posts ORDER BY created DESC")
if err != nil {
	panic(err.Error())
}

func dbInit(){

}

func (m Mysql)dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := m.user
	dbPass := m.password
	dbName := m.name
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func (m Mysql) getPosts {
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
}
