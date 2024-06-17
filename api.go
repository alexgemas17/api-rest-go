package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	address string
	store   Store
}

func NewApiServer(address string, store Store) *ApiServer {
	return &ApiServer{address: address, store: store}
}

func (s *ApiServer) Serve() {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	TasksService := NewTasksService(s.store)
	TasksService.RegisterRoutes(subRouter)

	log.Println("Started api serve at", s.address)
	log.Fatal(http.ListenAndServe(s.address, subRouter))
}
