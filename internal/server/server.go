package server

import (
	"encoding/json"
	"log"
	"net/http"

	db "github.com/Myronarty/Lab_Go/db/sqlc"
	"github.com/Myronarty/Lab_Go/internal/server/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
	store  db.Store
}

func NewServer(store db.Store) *Server {
	s := &Server{
		router: mux.NewRouter(),
		store:  store,
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.router.HandleFunc("/health", s.handleHealth).Methods("GET")

	// Create kogut handler
	kogutHandler := handlers.NewKogutHandler(s.store)

	// Kogut routes
	s.router.HandleFunc("/koguts", kogutHandler.GetAllKoguts).Methods("GET")
	s.router.HandleFunc("/koguts", kogutHandler.CreateKogut).Methods("POST")
	s.router.HandleFunc("/koguts/{id}", kogutHandler.GetKogut).Methods("GET")
	s.router.HandleFunc("/koguts/{id}", kogutHandler.UpdateKogut).Methods("PUT")
	s.router.HandleFunc("/koguts/{id}", kogutHandler.DeleteKogut).Methods("DELETE")
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (s *Server) Run(port string) {
	log.Fatal(http.ListenAndServe(port, s.router))
}
