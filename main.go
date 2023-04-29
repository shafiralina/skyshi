package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"skyshi/controllers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/activity-groups", controllers.CreateActivities).Methods("POST")
	router.HandleFunc("/api/activity-groups/{id}", controllers.UpdateActivities).Methods("PATCH")
	router.HandleFunc("/api/activity-groups", controllers.GetAllActivities).Methods("GET")
	router.HandleFunc("/api/activity-groups/{id}", controllers.GetActivities).Methods("GET")
	router.HandleFunc("/api/activity-groups/{id}", controllers.DeleteActivities).Methods("DELETE")
	router.HandleFunc("/api/todo-items", controllers.CreateTodos).Methods("POST")
	router.HandleFunc("/api/todo-items/{id}", controllers.UpdateTodos).Methods("PATCH")
	router.HandleFunc("/api/todo-items", controllers.GetAllTodos).Methods("GET")
	router.HandleFunc("/api/todo-items/{id}", controllers.GetTodos).Methods("GET")
	router.HandleFunc("/api/todo-items/{id}", controllers.DeleteTodos).Methods("DELETE")

	port := "3030"
	fmt.Println("Started at:", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
