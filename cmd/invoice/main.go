package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

type CardTransaction struct {
	CardNumber  string
	Description string
	Amount      int
	Currency    string
	Date        string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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
		dsn := fmt.Sprintf(
			`host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Sao_Paulo`,
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		)

		db, err := sql.Open(`postgres`, dsn)
		if err != nil {
			log.Printf("Failed to connect on database: %v\n", err)
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{
					`error`: `An error ocurred while trying to get invoice`,
				},
			)
		}
		defer db.Close()

		rows, err := db.Query(
			`SELECT * FROM cards_transaction WHERE card_number = $1`,
			c.Param(`cardNumber`),
		)
		if err != nil {
			log.Printf("Failed to get data from database: %v\n", err)
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{
					`error`: `An error ocurred while trying to get invoice`,
				},
			)
		}

		var cardsTransaction []CardTransaction
		for rows.Next() {
			var cardTransaction CardTransaction
			err := rows.Scan(
				&cardTransaction.CardNumber,
				&cardTransaction.Description,
				&cardTransaction.Amount,
				&cardTransaction.Currency,
				&cardTransaction.Date,
			)
			if err != nil {
				log.Printf("Failed to parse data from database: %v\n", err)
				return c.JSON(
					http.StatusInternalServerError,
					map[string]string{
						`error`: `An error ocurred while trying to get invoice`,
					},
				)
			}
			cardsTransaction = append(cardsTransaction, cardTransaction)
		}

		var total = 0
		for _, transaction := range cardsTransaction {
			total += transaction.Amount
		}

		return c.JSON(http.StatusOK, map[string]int{`total`: total})
	})

	e.Logger.Fatal(e.Start(`:3000`))
}
