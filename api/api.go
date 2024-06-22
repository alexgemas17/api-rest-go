package api

import (
	"log"
	"net/http"

	"github.com/alexgemas17/api-rest-go/api/store"
	"github.com/alexgemas17/api-rest-go/api/tasks"
	userService "github.com/alexgemas17/api-rest-go/api/users"
	"github.com/gorilla/mux"
)

type ApiServer struct {
	address string
	store   store.Store
}

func NewApiServer(address string, store store.Store) *ApiServer {
	return &ApiServer{address: address, store: store}
}

func (s *ApiServer) Serve() {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userService := userService.NewUserService(s.store)
	userService.RegisterRoutes(subRouter)

	tasksService := tasks.NewTasksService(s.store)
	tasksService.RegisterRoutes(subRouter)

	log.Println("Started api serve at", s.address)
	log.Fatal(http.ListenAndServe(s.address, subRouter))
}
