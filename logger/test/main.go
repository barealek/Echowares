package main

import (
	"errors"
	"net/http"

	"github.com/barealek/echowares/echologger"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Use(echologger.New())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/error", func(c echo.Context) error {
		return errors.New("Der skete en fejl")
	})

	e.GET("/redirect", func(c echo.Context) error {
		return c.Redirect(http.StatusPermanentRedirect, "/")
	})

	e.Logger.Fatal(e.Start("0.0.0.0:52333"))
}
