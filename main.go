package main

import (
	controller "dummige/controllers"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	e.GET("/", controller.Get)

	e.GET("/*", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "/")
	})

	port := "localhost:1323"
	remote := os.Getenv("PORT")

	if remote != "" {
		port = ":" + remote
	}

	e.Logger.Fatal(e.Start(port))

}