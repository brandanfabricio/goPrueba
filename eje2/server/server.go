package server

import "net/http"

type Country struct {
	Name     string
	Lenguaje string
}

var countries []*Country = []*Country{}

func New(addr string) *http.Server {
initRouter()

	return &http.Server{
		Addr: addr,

	}
}