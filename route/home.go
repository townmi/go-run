package route

import (
	"net/http"
	"fmt"
)

type Route interface {

}

func GetHome(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Fprintf(w, "Hello World")
}

func PostHome(w http.ResponseWriter, r *http.Request)  {

}