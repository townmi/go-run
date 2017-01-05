package main

import (
	"net/http"
	"fmt"
	"log"
	"bytes"
	_ "os"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"path"
)

var (
	tmpl []byte
	tmplpath string
)


func main() {

	tmplpath = "/Users/harry/golang/src/go-run/views/index.html"

	fmt.Println(path.IsAbs(tmplpath))

	buff, err := ioutil.ReadFile(tmplpath)

	if err != nil {
		panic("open file failed!")
	}

	tmpl = buff

	data := struct {
		Title string
		Items []string
	}{
		Title: "My page",
		Items: []string{
			"My photos",
			"My blog",
		},
	}

	t, err := template.New("index").Parse(string(buff));

	var b bytes.Buffer

	err = t.Execute(&b, data)

	tmpl = b.Bytes()

	if err != nil {
		panic("open file failed!")
	}

	r := mux.NewRouter()

	r.Host("www.example.com")

	// Routes consist of a path and a handler function.
	r.HandleFunc("/", index).Methods("GET")

	r.HandleFunc("/{cat}", YourHandler2)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}

func YourHandler2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write(tmpl)
}

