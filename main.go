package main

import (
	"net/http"

	"github.com/gorilla/mux" // Your imported package
)

func yourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func difHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Different Gorilla!\n"))
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	//r.HandleFunc("/gorilla", yourHandler)
	//r.HandleFunc("/new", difHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	// Bind to a port and pass our router in
	panic(http.ListenAndServe(":8000", r))
}
