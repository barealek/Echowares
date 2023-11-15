package main

import (
	"errors"
	"net/http"
	"time"

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

	e.GET("/waitalittle", func(c echo.Context) error {
		time.Sleep(2 * time.Millisecond)
		return c.String(http.StatusCreated, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":3000"))
}
