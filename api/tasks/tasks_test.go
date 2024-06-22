package tasks

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"github.com/alexgemas17/api-rest-go/api/store"
	"github.com/alexgemas17/api-rest-go/api/types"
)

func runMockRequest(payload *types.Task, t *testing.T, s *TasksService) *httptest.ResponseRecorder {
	// Generate the fake body request
	b, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}

	// Make fake request
	req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(b))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()

	// handle it with a mocked service
	router.HandleFunc("/tasks", s.handleCreateTasks)

	// serve the request
	router.ServeHTTP(rr, req)

	return rr
}

func TestCreateTask(t *testing.T) {
	ms := &store.MockStore{}
	service := NewTasksService(ms)

	t.Run("should return an error if name is empty", func(t *testing.T) {
		payload := &types.Task{
			Name: "",
		}

		rr := runMockRequest(payload, t, service)

		// double check just in case the request was successful
		if rr.Code != http.StatusBadRequest {
			t.Error("invalid status code, it should fail")
		}
	})

	t.Run("should creaste a Task", func(t *testing.T) {
		payload := &types.Task{
			Name:       "Test name",
			ProjectID:  1,
			AssignedTo: 2,
		}

		rr := runMockRequest(payload, t, service)

		// double check just in case the request was successful
		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})
}

func TestGetTask(t *testing.T) {
	ms := &store.MockStore{}
	service := NewTasksService(ms)

	t.Run("should return the task", func(t *testing.T) {
		// Make fake request
		req, err := http.NewRequest(http.MethodGet, "/tasks/42", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		// handle it with a mocked service
		router.HandleFunc("/tasks/{id}", service.handleGetTask)

		// serve the request
		router.ServeHTTP(rr, req)

		// double check just in case the request was successful
		if rr.Code != http.StatusOK {
			t.Error("invalid status code, it should fail")
		}
	})
}
