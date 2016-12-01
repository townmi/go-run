package rest

import (
	"net/http"
	"fmt"
)

type routeInfo struct {
	url string
	method string
	handler HandlerFunc
}

type HandlerFunc func(http.ResponseWriter, *http.Request)


type Init struct {
	Route map[string]HandlerFunc
}

var R = &Init{make(map[string]HandlerFunc)}

func (R *Init) Map() {

}

func (R *Init) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(len(R.Route))
	if r.URL.Path == "/" {
		fmt.Println("123")
		return
	}
	http.NotFound(w, r)
	return
}

/**
    register router
 */
func (R *Init) Post(url string, h HandlerFunc) {
	rInfo := routeInfo{url, "POST", h}
	R.Route[url] = rInfo
}
