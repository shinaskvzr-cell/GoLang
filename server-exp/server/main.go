package main

import (
	"fmt"
	"log"
	"net/http"
)

func PingRequest(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello! You requested: %s", r.URL.Path)

}

func main() {

	http.HandleFunc("/", PingRequest)

	log.Println(http.ListenAndServe(":8080", nil))

}