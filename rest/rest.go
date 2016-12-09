package rest

import (
	"net/http"
	"strings"
	"log"
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

	key := strings.ToUpper(r.Method) + strings.ToUpper(r.URL.Path)

	if url := R.Route[key].url; len(url) != 0 {
		R.Route[key].handler(w, r);
		return
	}
	http.NotFound(w, r)
	return
}

/**
    register router
 */
func (R *Init) Post(url string, h HandlerFunc) {

	key := "POST" + strings.ToUpper(url)

	rInfo := routeInfo{url, "POST", h}

	R.Route[key] = rInfo
	log.Fatal("Register router:" )
}

func (R *Init) Get(url string, h HandlerFunc) {

	key := "GET" + strings.ToUpper(url)

	rInfo := routeInfo{url, "GET", h}
	R.Route[key] = rInfo
}

func (R *Init) Put(url string, h HandlerFunc) {

	key := "PUT" + strings.ToUpper(url)

	rInfo := routeInfo{url, "PUT", h}
	R.Route[key] = rInfo
}

func (R *Init) Delete(url string, h HandlerFunc) {

	key := "DELETE" + strings.ToUpper(url)

	rInfo := routeInfo{url, "DELETE", h}
	R.Route[key] = rInfo
}
