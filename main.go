package main

import (
	"log"
	"net/http"

	"github.com/tomihaapalainen/go-htmx-form/handler"
)

func main() {
	r := http.NewServeMux()

	r.Handle("/", handler.HandleIndex())
	r.HandleFunc("/firstName", handler.HandlePostFirstName)
	r.HandleFunc("/lastName", handler.HandlePostLastName)
	r.HandleFunc("/email", handler.HandlePostEmail)
	r.HandleFunc("/password", handler.HandlePostPassword)

	log.Fatal(http.ListenAndServe(":8080", r))
}
