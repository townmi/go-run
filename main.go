package main

import (
	"net/http"
	"go-run/config"
	"log"
	"go-run/route"
	"github.com/gorilla/mux"
)

func main() {

	// new mux Router
	r := mux.NewRouter()
	//
	r.Host("www.example.com")

	// GET ROUTES MAP
	r.HandleFunc("/", route.GetHome).Methods("GET")
	r.HandleFunc("/search", route.GetSearch).Methods("GET")


	r.HandleFunc("/email", route.SendEmail).Methods("GET")

	// POST ROUTES MAP
	r.HandleFunc("/search", route.PostSearch).Methods("POST")

	// BIND PORT TO SERVER
	log.Fatal(http.ListenAndServe(":"+config.Env.PORT, r))

}
