package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello From Service B!\n"))
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/test", HandleRoot)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":80", r))
}
