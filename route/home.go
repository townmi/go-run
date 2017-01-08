package route

import (
	"net/http"
	"html/template"
	"bytes"
	"io/ioutil"
	"go-run/config"
)

var (
	html []byte
)

type MapRoute interface {
}

func GetHome(w http.ResponseWriter, r *http.Request) {

	// get index view template
	viewPath := config.Env.PATH + "views/index.html"
	buff, err := ioutil.ReadFile(viewPath)
	config.CheckError(err)

	// set data model, send to view
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

	// new template with index template
	t, err := template.New("index").Parse(string(buff))
	config.CheckError(err)

	// write data to view
	var b bytes.Buffer
	err = t.Execute(&b, data)
	config.CheckError(err)

	// send document type of bytes to client
	html = b.Bytes()
	w.Write(html)
}
