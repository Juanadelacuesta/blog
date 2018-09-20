package main

import (
	"fmt"
	"reflect"
	"time"
)

type Post struct {
	Id    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
	Date  string `json:"date,omitempty"`
}

type Posts []Post

func (p Post) FormatDate(date time.Time) (s string) {
	s = fmt.Sprintf("%s %dth %d", date.Month(), date.Day(), date.Year())
	return
}

func (p *Post) IsEmpty() bool {
	return reflect.DeepEqual(p, Post{})
}
