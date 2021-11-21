package server

import "net/http"

type City struct {
	Name    string
	Country string
}

var cities []*City = []*City{}

func New(addr string) *http.Server {
	initRoutes()
	return &http.Server{
		Addr: addr,
	}
}
