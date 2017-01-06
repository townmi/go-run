package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"go-run/route"
	"go-run/config"
)

func main() {

	r := mux.NewRouter()

	r.Host("www.example.com")

	// GET ROUTES MAP
	r.HandleFunc("/", route.GetHome).Methods("GET")
	r.HandleFunc("/search", route.GetSearch).Methods("GET")

	// POST ROUTES MAP

	// BIND PORT TO SERVER
	log.Fatal(http.ListenAndServe(":"+config.Env.PORT, r))
}
