package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jene4kabezdar/biocad/apiserver/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/v1/page/{number}", handlers.PaginatonHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}