package main

import (
	"first-go/db"
	"first-go/handler"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Worlds")
}

func handleRequests() {
	//using gorilla mux
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", helloWorld).Methods("GET")

	//users endpoint
	r.HandleFunc("/users", handler.AllUsers).Methods("GET")
	r.HandleFunc("/user", handler.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", handler.UpdateUser).Methods("PATCH")
	r.HandleFunc("/user/{id}", handler.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", r))
}

func main() {
	fmt.Println("Go run on port 8081")

	//db operation
	db.Migrate()

	//run the func
	handleRequests()
}
