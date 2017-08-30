package route

import (
	"bytes"
	"net/http"

	"github.com/go-ini/ini"
	"github.com/labstack/echo"
)

var (
	venuesDbLink string
)

// var data map[string]model.User

// VenuesGraphql graphql for api
func VenuesGraphql(c echo.Context) error {
	// session, err := mgo.Dial(venuesDbLink)
	// if err != nil {
	// 	panic(err)
	// }
	// session.Close()

	// result := executeQuery(c.QueryParam("query"), schema)

	return c.JSONPretty(http.StatusOK, 1, "  ")
}

func init() {
	cfg, err := ini.Load("conf.ini")
	if err != nil {
		panic(err)
	}
	USERNAME := cfg.Section("VENUESDATABASE").Key("USERNAME").String()
	PASSWORD := cfg.Section("VENUESDATABASE").Key("PASSWORD").String()
	HOSTNAME := cfg.Section("VENUESDATABASE").Key("HOSTNAME").String()
	DATABASE := cfg.Section("VENUESDATABASE").Key("DATABASE").String()
	PORT := cfg.Section("VENUESDATABASE").Key("PORT").String()

	b := bytes.Buffer{}
	b.WriteString("mongodb://")
	b.WriteString(USERNAME)
	b.WriteString(":")
	b.WriteString(PASSWORD)
	b.WriteString("@")
	b.WriteString(HOSTNAME)
	b.WriteString(":")
	b.WriteString(PORT)
	b.WriteString("/")
	b.WriteString(DATABASE)
	venuesDbLink = b.String()
}
