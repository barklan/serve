package main

import (
	"log"
	"net/http"
)

const (
	port      = "9999"
	directory = "/public"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(directory)))
	log.Printf("Serving on HTTP port: %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
