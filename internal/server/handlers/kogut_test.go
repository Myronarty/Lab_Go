package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	db "github.com/Myronarty/Lab_Go/db/sqlc"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgtype"
)

type DummyStore struct{}

func (ds *DummyStore) CreateKogut(ctx context.Context, arg db.CreateKogutParams) (db.Kogut, error) {
	return db.Kogut{
		ID:   1,
		Name: arg.Name,
		Age:  arg.Age,
		Sex:  arg.Sex,
	}, nil
}

func (ds *DummyStore) GetKogut(ctx context.Context, id int32) (db.Kogut, error) {
	return db.Kogut{
		ID:   id,
		Name: "TestKogut",
		Age:  pgtype.Int4{Int32: 5, Valid: true},
		Sex:  true,
	}, nil
}

func (ds *DummyStore) UpdateKogut(ctx context.Context, arg db.UpdateKogutParams) (db.Kogut, error) {
	return db.Kogut{
		ID:   arg.ID,
		Name: arg.Name,
		Age:  arg.Age,
		Sex:  arg.Sex,
	}, nil
}

func (ds *DummyStore) DeleteKogut(ctx context.Context, id int32) error {
	return nil
}

func (ds *DummyStore) GetAllKoguts(ctx context.Context) ([]db.Kogut, error) {
	return []db.Kogut{
		{ID: 1, Name: "Kogut1", Age: pgtype.Int4{Int32: 5, Valid: true}, Sex: true},
		{ID: 2, Name: "Kogut2", Age: pgtype.Int4{Int32: 3, Valid: true}, Sex: false},
	}, nil
}

func TestCreateKogut(t *testing.T) {
	dummyStore := &DummyStore{}
	handler := NewKogutHandler(dummyStore)

	payload := map[string]interface{}{
		"name": "TestKogut",
		"age":  5,
		"sex":  true,
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest("POST", "/koguts", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.CreateKogut(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("Expected status 201 Created, got %v", rec.Code)
	}

	var result db.Kogut
	json.NewDecoder(rec.Body).Decode(&result)

	if result.Name != "TestKogut" {
		t.Errorf("Expected name 'TestKogut', got %v", result.Name)
	}
}

func TestGetKogut(t *testing.T) {
	dummyStore := &DummyStore{}
	handler := NewKogutHandler(dummyStore)

	req := httptest.NewRequest("GET", "/koguts/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rec := httptest.NewRecorder()

	handler.GetKogut(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", rec.Code)
	}
}

func TestUpdateKogut(t *testing.T) {
	dummyStore := &DummyStore{}
	handler := NewKogutHandler(dummyStore)

	payload := map[string]interface{}{
		"name": "UpdatedName",
		"age":  6,
		"sex":  false,
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest("PUT", "/koguts/1", bytes.NewReader(body))
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rec := httptest.NewRecorder()

	handler.UpdateKogut(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", rec.Code)
	}
}

func TestDeleteKogut(t *testing.T) {
	dummyStore := &DummyStore{}
	handler := NewKogutHandler(dummyStore)

	req := httptest.NewRequest("DELETE", "/koguts/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rec := httptest.NewRecorder()

	handler.DeleteKogut(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Errorf("Expected 204 No Content, got %v", rec.Code)
	}
}
