package vmw


import (
	"github.com/sclevine/agouti"
)


func Ex(){

	driver := agouti.Selenium()

	driver.Start()

	page, _ := driver.NewPage(agouti.Browser("chrome"))

	page.Size(200, 300)

	page.Navigate("http://lp.vipabc.com/program/linkage_page/newjsapitest/index.html")

	var number int

	page.RunScript("alert(1);", map[string]interface{}{"test": 100}, &number)

}