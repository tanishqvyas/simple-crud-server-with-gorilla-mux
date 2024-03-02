package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
)

func Create(w http.ResponseWriter, r *http.Request) {
	// Read the request body
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error reading request body", http.StatusInternalServerError)
        return
    }

    defer r.Body.Close()

    // Print the request body
    fmt.Fprintf(w, "Received body: %s", body)
}

func Read(w http.ResponseWriter, r *http.Request) {
	// Here you would handle the read logic, for simplicity we'll just return a message
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "Reading item with ID %s", id)
}

func Update(w http.ResponseWriter, r *http.Request) {
	// Here you would handle the update logic, for simplicity we'll just return a message
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "Updating item with ID %s", id)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	// Here you would handle the delete logic, for simplicity we'll just return a message
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "Deleting item with ID %s", id)
}
