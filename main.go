package main

import (
	"net/http"

	"./route"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	e.GET("/user/:id", route.GetUserInfo)
	e.PUT("/user/:id", route.PutUserInfo)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
