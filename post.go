package main

import (
	"fmt"
	"time"
)

type Post struct {
	Id    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
	Date  string `json:"date,omitempty"`
}

type Posts []Post

func (Post) FormatDate(date time.Time) (s string) {
	s = fmt.Sprintf("%s %dth %d", date.Month(), date.Day(), date.Year())
	return
}
