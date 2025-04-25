package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	db "github.com/nochzato/example/db/sqlc"
)

type Server struct {
	router *chi.Mux
	store  db.Store
}

func NewServer(store db.Store) *Server {
	s := &Server{
		router: chi.NewRouter(),
		store:  store,
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.router.Get("/health", s.handleHealth)
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (s *Server) Run(port string) {
	log.Fatal(http.ListenAndServe(port, s.router))
}
