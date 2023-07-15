package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: "method=${method}, uri=${uri}, status=${status}\n"}))
	e.Logger.Fatal(e.Start(`:3000`))
}
