package store

import (
	"database/sql"

	"github.com/alexgemas17/api-rest-go/api/types"
)

type Store interface {
	// Users
	CreateUser(u *types.User) (*types.User, error)
	GetUserById(id string) (*types.User, error)
	GetUserByEmail(email string) (*types.User, error)

	CreateTask(t *types.Task) (*types.Task, error)
	GetTask(id string) (*types.Task, error)
}

type Repository struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (s *Repository) CreateUser(u *types.User) (*types.User, error) {
	rows, err := s.db.Exec("INSERT INTO users (email, firstName, lastName, password) VALUES (?, ?, ?, ?)", u.Email, u.FirstName, u.LastName, u.Password)
	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	u.ID = id
	return u, nil
}

func (s *Repository) GetUserById(id string) (*types.User, error) {
	var u types.User
	err := s.db.QueryRow("SELECT id, email, firstName, lastName, createdAt FROM users WHERE id = ?", id).Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.CreatedAt)
	return &u, err
}

func (s *Repository) GetUserByEmail(email string) (*types.User, error) {
	var u types.User
	err := s.db.QueryRow("SELECT id, email, firstName, lastName, password, createdAt FROM users WHERE email = ?", email).Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.Password, &u.CreatedAt)
	return &u, err
}

func (s *Repository) CreateTask(t *types.Task) (*types.Task, error) {
	rows, err := s.db.Exec("INSERT INTO tasks (name, status, project_id, assigned_to) VALUES (?, ?, ?, ?)", t.Name, t.Status, t.ProjectID, t.AssignedTo)

	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	t.ID = id
	return t, nil
}

func (s *Repository) GetTask(id string) (*types.Task, error) {
	var t types.Task
	err := s.db.QueryRow("SELECT id, name, status, project_id, assigned_to, createdAt FROM tasks WHERE id = ?", id).Scan(&t.ID, &t.Name, &t.Status, &t.ProjectID, &t.AssignedTo, &t.CreatedAt)
	return &t, err
}

/* Mock types */
type MockStore struct{}

func (m *MockStore) CreateUser(u *types.User) (*types.User, error) {
	return &types.User{}, nil
}

func (m *MockStore) GetUserById(id string) (*types.User, error) {
	return &types.User{}, nil
}

func (m *MockStore) GetUserByEmail(email string) (*types.User, error) {
	return &types.User{}, nil
}

func (m *MockStore) CreateTask(t *types.Task) (*types.Task, error) {
	return &types.Task{}, nil
}

func (m *MockStore) GetTask(id string) (*types.Task, error) {
	return &types.Task{}, nil
}
