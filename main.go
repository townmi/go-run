package main

 import (
 	"log"
 	"net/http"
 	"go-run/rest"
 	"go-run/route"
 )


 func main(){
 	var app = *rest.R

 	err := http.ListenAndServe(":9090", rest.R) //设置监听的端口

 	if err != nil {
 		log.Fatal("ListenAndServe: ", err)
 	}

 	app.Get("/", route.GetHome);

 	app.Get("/2", route.GetHome);

 	app.Post("/", route.PostHome);

 }