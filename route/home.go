package route

import (
	"net/http"
	"fmt"
	"html/template"
	"bytes"
	"io/ioutil"
	"go-run/config"
)

var (
	tmpl []byte
)

type Route interface {
}

func GetHome(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("path: %v", "in")

	tmplpath := config.Env.PATH + "views/index.html"

	fmt.Printf("path: %v\n", tmplpath)

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

	t, err := template.New("index").Parse(string(buff))
	var b bytes.Buffer

	err = t.Execute(&b, data)

	tmpl = b.Bytes()

	if err != nil {
		panic("open file failed!")
	}

	w.Write(tmpl)
}

func PostHome(w http.ResponseWriter, r *http.Request) {

}
