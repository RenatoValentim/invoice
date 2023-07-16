package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(
		middleware.LoggerWithConfig(
			middleware.LoggerConfig{
				Format: "method=${method}, uri=${uri}, status=${status}\n",
			},
		),
	)

	e.GET(`/currencies`, func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]int{`usd`: 3})
	})

	e.Logger.Fatal(e.Start(`:3001`))
}
