package route

import (
	"net/http"
)

func init() {

}

func GetHouse(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(string("hello word!")))
}