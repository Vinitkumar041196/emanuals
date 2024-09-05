package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	log.Println("Got request on ", r.URL.Path)
	fmt.Fprintf(w, "<h1>Welcome to E-Manuals Home Page</h1>")
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.ListenAndServe(":3000", nil)
}
