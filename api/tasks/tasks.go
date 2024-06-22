package tasks

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/alexgemas17/api-rest-go/api/auth"
	"github.com/alexgemas17/api-rest-go/api/store"
	"github.com/alexgemas17/api-rest-go/api/types"
	"github.com/alexgemas17/api-rest-go/api/utils"
	"github.com/gorilla/mux"
)

var errNameRequired = errors.New("name is required")
var errProjectIdRequired = errors.New("projectID is required")
var errUserIdRequired = errors.New("userID is required")

type TasksService struct {
	store store.Store
}

func NewTasksService(s store.Store) *TasksService {
	return &TasksService{store: s}
}

func (s *TasksService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", auth.WithJWTAuth(s.handleCreateTasks, s.store)).Methods("POST")
	r.HandleFunc("/tasks/{id}", auth.WithJWTAuth(s.handleGetTask, s.store)).Methods("GET")
}

func (s *TasksService) handleCreateTasks(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, types.ErrorResponse{Error: "Invalid request payload"})
		return
	}

	defer r.Body.Close()

	var task *types.Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, types.ErrorResponse{Error: "Invalid request payload"})
		return
	}

	if err := validateTaskPayload(task); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, types.ErrorResponse{Error: "Invalid payload: " + err.Error()})
		return
	}

	t, err := s.store.CreateTask(task)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, types.ErrorResponse{Error: "Error creating task"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, t)
}

func (s *TasksService) handleGetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		utils.WriteJSON(w, http.StatusBadRequest, types.ErrorResponse{Error: "ID is required"})
	}

	t, err := s.store.GetTask(id)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, types.ErrorResponse{Error: "Task not found"})
	}

	utils.WriteJSON(w, http.StatusOK, t)
}

func validateTaskPayload(task *types.Task) error {
	if task.Name == "" {
		return errNameRequired
	}

	if task.ProjectID == 0 {
		return errProjectIdRequired
	}

	if task.AssignedTo == 0 {
		return errUserIdRequired
	}

	return nil
}
