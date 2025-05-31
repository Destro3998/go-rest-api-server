package router

import (
	"github.com/Destro3998/go-rest-api/handlers"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
	return r
}
