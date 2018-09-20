package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func (b *Blog) GetRoutes() Routes {
	return Routes{
		Route{
			"Index",
			"GET",
			"/",
			b.Index,
		},
		Route{
			"Show",
			"GET",
			"/post/{id}",
			b.Show,
		},
		Route{
			"Search",
			"GET",
			"/search",
			b.Search,
		},
	}
}
