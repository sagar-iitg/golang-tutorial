package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{}", GetUserById).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9000", r))

}
func main() {
	InitialMigration()
	initializeRoutes()

}
