package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "This is the entry point for the API")
	})
	e.Logger.Fatal(e.Start(":8001"))
}
