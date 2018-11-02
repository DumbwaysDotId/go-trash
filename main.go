package main

import (
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
	log.Fatal(http.ListenAndServe(":8081", r))
}

func main() {
	fmt.Println("Go run on port 8081")

	//run the func
	handleRequests()
}
