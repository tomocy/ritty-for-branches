package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func New() *Server {
	return new(Server)
}

type Server struct {
	router chi.Router
}

func (s *Server) ListenAndServe(addr string) error {
	log.Printf("listen and serve on %s\n", addr)
	return http.ListenAndServe(addr, s.router)
}
