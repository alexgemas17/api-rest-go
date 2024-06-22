package types

import "time"

type ErrorResponse struct {
	Error string `json:"error"`
}

type Task struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Status     string    `json:"status"`
	ProjectID  int64     `json:"ProjectID"`
	AssignedTo int64     `json:"AssignedTo"`
	CreatedAt  time.Time `json:"createdat"`
}

type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdat"`
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
