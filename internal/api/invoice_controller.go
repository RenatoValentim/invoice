package api

import (
	"invoice/internal/usecase"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InvoiceController(calculateInvoice usecase.CalculateInvoice) {
	e := echo.New()

	e.Use(
		middleware.LoggerWithConfig(
			middleware.LoggerConfig{
				Format: "method=${method}, uri=${uri}, status=${status}\n",
			},
		),
	)

	e.GET(`/health`, func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{`status`: `ok`})
	})

	e.GET(`/cards/:cardNumber/invoices`, func(c echo.Context) error {
		total, err := calculateInvoice.Execute(c.Param(`cardNumber`))
		if err != nil {
			log.Printf("Failed when calculate invoice: %v\n", err)

			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{
					`error`: `An error ocurred while trying to get invoice`,
				},
			)
		}

		return c.JSON(http.StatusOK, map[string]float64{`total`: total})
	})

	e.Logger.Fatal(e.Start(`:3000`))
}
