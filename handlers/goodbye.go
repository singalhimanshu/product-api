package handlers

import (
	"log"
	"net/http"
)

type GoodBye struct {
	l log.Logger
}

func NewGoodBye(l *log.Logger) *GoodBye {
	return &GoodBye{}
}

func (g *GoodBye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GoodBye"))
}
