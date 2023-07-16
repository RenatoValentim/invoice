package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

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

type Currency struct {
	USD int `json:"usd"`
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

		currentDate := time.Now()
		month := currentDate.Month()
		year := currentDate.Year()

		response, err := http.Get(`http://172.17.0.1:3001/currencies`)
		if err != nil {
			log.Printf("Failed when get currency: %v\n", err)
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{
					`error`: `An error ocurred while trying to get invoice`,
				},
			)
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Printf("Failed when read response.body: %v\n", err)
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{
					`error`: `An error ocurred while trying to get invoice`,
				},
			)
		}

		var currency Currency
		err = json.Unmarshal(body, &currency)
		if err != nil {
			log.Printf("Failed when unmarshal currency: %v\n", err)
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{
					`error`: `An error ocurred while trying to get invoice`,
				},
			)
		}

		rows, err := db.Query(`
			SELECT * FROM cards_transaction
			WHERE card_number = $1
			AND EXTRACT(month from date) = $2
			AND EXTRACT(year from date) = $3`,
			c.Param(`cardNumber`),
			int(month),
			year,
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
			if transaction.Currency == `BRL` {
				total += transaction.Amount
			} else {
				total += transaction.Amount * currency.USD
			}
		}

		return c.JSON(http.StatusOK, map[string]int{`total`: total})
	})

	e.Logger.Fatal(e.Start(`:3000`))
}
