package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	db "github.com/Myronarty/Lab_Go/db/sqlc"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgtype"
)

type KogutHandler struct {
	store db.Store
}

func NewKogutHandler(store db.Store) *KogutHandler {
	return &KogutHandler{store: store}
}

// CreateKogut handles POST /koguts
func (h *KogutHandler) CreateKogut(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
		Age  *int32 `json:"age,omitempty"`
		Sex  bool   `json:"sex"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	// Handle nullable age
	var age pgtype.Int4
	if req.Age != nil {
		age = pgtype.Int4{Int32: *req.Age, Valid: true}
	} else {
		age = pgtype.Int4{Valid: false}
	}

	// Use your existing sqlc generated function
	kogut, err := h.store.CreateKogut(r.Context(), db.CreateKogutParams{
		Name: req.Name,
		Age:  age,
		Sex:  req.Sex,
	})
	if err != nil {
		http.Error(w, "Failed to create kogut", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(kogut)
}

// GetKogut handles GET /koguts/{id}
func (h *KogutHandler) GetKogut(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Use your existing sqlc generated function
	kogut, err := h.store.GetKogut(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Kogut not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kogut)
}

// GetAllKoguts handles GET /koguts
func (h *KogutHandler) GetAllKoguts(w http.ResponseWriter, r *http.Request) {
	// Use your existing sqlc generated function
	koguts, err := h.store.GetAllKoguts(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch koguts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(koguts)
}

// UpdateKogut handles PUT /koguts/{id}
func (h *KogutHandler) UpdateKogut(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var req struct {
		Name string `json:"name"`
		Age  *int32 `json:"age,omitempty"`
		Sex  bool   `json:"sex"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	// Handle nullable age
	var age pgtype.Int4
	if req.Age != nil {
		age = pgtype.Int4{Int32: *req.Age, Valid: true}
	} else {
		age = pgtype.Int4{Valid: false}
	}

	// Use your existing sqlc generated function
	kogut, err := h.store.UpdateKogut(r.Context(), db.UpdateKogutParams{
		ID:   int32(id),
		Name: req.Name,
		Age:  age,
		Sex:  req.Sex,
	})
	if err != nil {
		http.Error(w, "Failed to update kogut", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kogut)
}

// DeleteKogut handles DELETE /koguts/{id}
func (h *KogutHandler) DeleteKogut(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Use your existing sqlc generated function
	err = h.store.DeleteKogut(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Failed to delete kogut", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
