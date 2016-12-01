package rest

import (
	"net/http"
)

type route struct {
	url string
	method string
	handler HandlerFunc
}

type HandlerFunc func(http.ResponseWriter, *http.Request)


type Init struct {
	route map[int]struct{}
}

func (R *Init) Map() {

}

func (R *Init) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		return
	}
	http.NotFound(w, r)
	return
}


/**
    register router
 */
func (R *Init) Post(r, h HandlerFunc) {
	R.route[len(R.route)] = route{url: "/", method:"POST", "handler": h}
}