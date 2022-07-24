package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(context echo.Context) error {
		app := context.Request().Header.Get("app")
		return context.String(http.StatusOK, app)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
