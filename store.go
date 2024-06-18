package main

import "database/sql"

type Store interface {
	// Users
	CreateUser(u *User) (*User, error)
	GetUserById(id string) (*User, error)
	GetUserByEmail(email string) (*User, error)

	CreateTask(t *Task) (*Task, error)
	GetTask(id string) (*Task, error)
}

type Repository struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (s *Repository) CreateUser(u *User) (*User, error) {
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

func (s *Repository) GetUserById(id string) (*User, error) {
	var u User
	err := s.db.QueryRow("SELECT id, email, firstName, lastName, createdAt FROM users WHERE id = ?", id).Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.CreatedAt)
	return &u, err
}

func (s *Repository) GetUserByEmail(email string) (*User, error) {
	var u User
	err := s.db.QueryRow("SELECT id, email, firstName, lastName, password, createdAt FROM users WHERE email = ?", email).Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.Password, &u.CreatedAt)
	return &u, err
}

func (s *Repository) CreateTask(t *Task) (*Task, error) {
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

func (s *Repository) GetTask(id string) (*Task, error) {
	var t Task
	err := s.db.QueryRow("SELECT id, name, status, project_id, assigned_to, createdAt FROM tasks WHERE id = ?", id).Scan(&t.ID, &t.Name, &t.Status, &t.ProjectID, &t.AssignedTo, &t.CreatedAt)
	return &t, err
}
