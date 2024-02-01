package main

import (
	"golang/database"
	"golang/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	database.ConnectWithDb()

	router := mux.NewRouter()

	router.HandleFunc("/users", handlers.HandlerGetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.HandlerGetUser).Methods("GET")
	router.HandleFunc("/users/create", handlers.HandlerCreateUser).Methods("POST")
	router.HandleFunc("/users/delete/{id}", handlers.HandlerDeleteUser).Methods("DELETE")
	router.HandleFunc("/users/update/{id}", handlers.HandlerUpdateUser).Methods("PUT")
	router.HandleFunc("/", handlers.HandlerHomePage).Methods("GET")

	http.ListenAndServe(":8080", router)

}
