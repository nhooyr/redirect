package main

import (
	"log"
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	u := r.URL.Query().Get("to")
	if u == "" {
		u = "http://google.com"
		log.Printf("redirecting %s to default http://google.com", r.RemoteAddr)
	} else {
		log.Printf("redirecting %s to %s", r.RemoteAddr, u)
	}
	http.Redirect(w, r, u, http.StatusFound)
}
