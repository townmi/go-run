package main

import (
	_ "github.com/astaxie/beego"
	"log"
	"net/http"
	"fmt"
	"go-run/rest"
)
func sayhelloName(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	fmt.Println(r.Form)

	fmt.Fprintf(w, "Hello World")
}

func main(){
	var app = *rest.R

	app.Post("/", sayhelloName);

	app.Post("/2", sayhelloName);

	err := http.ListenAndServe(":9090", rest.R) //设置监听的端口

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}