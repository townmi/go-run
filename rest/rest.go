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
	Route map[string]routeInfo
}

var R = &Init{make(map[string]routeInfo)}

func (R *Init) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if url := R.Route[r.URL.Path].url; len(url) != 0 {
		fmt.Println(R.Route[r.URL.Path])
		R.Route[r.URL.Path].handler(w, r);
		return
	}
	fmt.Println(r.URL.Path)
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

func (R *Init) Get(url string, h HandlerFunc) {
	rInfo := routeInfo{url, "GET", h}
	R.Route[url] = rInfo
}

func (R *Init) Put(url string, h HandlerFunc) {
	rInfo := routeInfo{url, "PUT", h}
	R.Route[url] = rInfo
}

func (R *Init) Delete(url string, h HandlerFunc) {
	rInfo := routeInfo{url, "DELETE", h}
	R.Route[url] = rInfo
}
