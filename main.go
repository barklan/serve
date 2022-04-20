package main

import (
	"log"
	"net/http"
)

const (
	port      = "8100"
	directory = "/public"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(directory)))
	log.Printf("Serving %s on HTTP port: %s\n", directory, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
