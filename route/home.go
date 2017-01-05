package route

import (
	"net/http"
	"fmt"
)

type Route interface {

}

func GetHome(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Form)
	fmt.Fprintf(w, "Hello World")
}

func PostHome(w http.ResponseWriter, r *http.Request) {

}