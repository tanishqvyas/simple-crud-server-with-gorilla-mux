package main

import (
    "github.com/tanishqvyas/simple-crud-server-with-gorilla-mux/handlers"
    "net/http"
    "github.com/gorilla/mux"
	"fmt"
)

func main() {
    r := mux.NewRouter()

    // Define routes
    r.HandleFunc("/items", handlers.Create).Methods("POST")
    r.HandleFunc("/items/{id}", handlers.Read).Methods("GET")
    r.HandleFunc("/items/{id}", handlers.Update).Methods("PUT")
    r.HandleFunc("/items/{id}", handlers.Delete).Methods("DELETE")

	fmt.Println("Listening atport: 8080")

    err := http.ListenAndServe("127.0.0.1:8080", r)

	if err != nil {
		fmt.Println("error occured while starting the server", err)
	}
}