package main

import (
	"net/http"
	"go-run/config"
	"log"
	"go-run/route"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
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
	r.HandleFunc("/stocklist", route.GetStockList).Methods("GET")
	r.HandleFunc("/stock", route.GetStock).Methods("GET")

	// POST ROUTES MAP
	r.HandleFunc("/search", route.PostSearch).Methods("POST")

	r.HandleFunc("/stocklist", route.ReFreshStockList).Methods("POST")
	r.HandleFunc("/stock", route.ReFreshStock).Methods("POST")

	// BIND PORT TO SERVER
	log.Fatal(http.ListenAndServe(":"+config.Env.PORT, handlers.CompressHandler(r)))

}
