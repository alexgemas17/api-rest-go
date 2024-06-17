package main

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
