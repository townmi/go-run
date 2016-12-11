package potato_test

import (
	"testing"
	_ "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	"fmt"
	"bytes"
)

//var (
//	Expect := gomega.Expect
//	Succeed := gomega.Succeed
//	HaveOccurred := gomega.HaveOccurred
//	HaveText := matchers.HaveText
//	HaveURL := matchers.HaveURL
//)



func TestUserLoginPrompt(t *testing.T) {

	RegisterTestingT(t)
	b := bytes.NewBuffer(make([]byte, 0))
	driver := agouti.Selenium()
	Expect(driver.Start()).To(Succeed())
	page, err := driver.NewPage(agouti.Browser("chrome"))
	Expect(err).NotTo(HaveOccurred())

	page.Size(200, 300)

	page.Navigate("http://lp.vipabc.com/program/linkage_page/newjsapitest/index.html")
	var number int

	page.RunScript("alert(1);", map[string]interface{}{"test": 100}, &number)

	//page.Find("#prompt")).To(HaveText("Please login!"))

	//t:= reflect.TypeOf(page);

	//fmt.Println(page.HTML())





	//fmt.Println("type:", reflect.TypeOf(ab))

	fmt.Fprintln(b)
	// Expect(driver.Stop()).To(Succeed()) // calls page.Destroy() automatically
}