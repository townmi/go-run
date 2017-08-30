package main

import (
	"net/http"

	"./route"
	"github.com/go-ini/ini"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	// user services
	// venues services
	e.GET("/user/graphql", route.UsersGraphql)
	e.GET("/user/:id", route.GetUserInfo)
	e.PUT("/user/:id", route.PutUserInfo)

	// venues services
	e.GET("/venues/graphql", route.VenuesGraphql)

	// Start server
	cfg, err := ini.Load("conf.ini")
	if err != nil {
		panic(err)
	}
	port := cfg.Section("SERVER").Key("PORT").String()

	e.Logger.Fatal(e.Start(":" + port))
}
