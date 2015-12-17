package main

import (
	"log"
	"net/http"
)

func main() {
	log.SetFlags(0)
	log.Fatal(http.ListenAndServe(":80", http.HandlerFunc(redirect)))
}
