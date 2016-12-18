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
		R.Route[key].handler(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

/**
 * [Post register router]
 * @Author   townmi
 * @DateTime 2016-12-13T22:42:56+0800
 * @url    {[string]} 	[description]
 * @h      {[HandlerFunc]}	[description]
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

/**
 * use method
 */

// import (
// 	"log"
// 	"net/http"
// 	"go-run/rest"
// 	"go-run/route"
// )


// func main(){
// 	var app = *rest.R

// 	err := http.ListenAndServe(":9090", rest.R) //设置监听的端口

// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}

// 	app.Get("/", route.GetHome);

// 	app.Get("/2", route.GetHome);

// 	app.Post("/", route.PostHome);

// }

