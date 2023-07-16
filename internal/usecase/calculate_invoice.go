package usecase

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type CardTransaction struct {
	CardNumber  string
	Description string
	Amount      float64
	Currency    string
	Date        string
}

type Currency struct {
	USD float64 `json:"usd"`
}

type calculateInvoice struct {
}

func NewCalculateInvoice() *calculateInvoice {
	return &calculateInvoice{}
}

func (c *calculateInvoice) Execute(cardNumber string) (float64, error) {
	dsn := fmt.Sprintf(
		`host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Sao_Paulo`,
		viper.GetString(`db_host`),
		viper.GetString(`db_user`),
		viper.GetString(`db_password`),
		viper.GetString(`db_name`),
	)

	db, err := sql.Open(`postgres`, dsn)
	if err != nil {
		log.Printf("Failed to connect on database: %v\n", err)
		return -1, err
	}
	defer db.Close()

	currentDate := time.Now()
	month := currentDate.Month()
	year := currentDate.Year()

	response, err := http.Get(`http://172.17.0.1:3001/currencies`)
	if err != nil {
		log.Printf("Failed when get currency: %v\n", err)
		return -1, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Failed when read response.body: %v\n", err)
		return -1, err
	}

	var currency Currency
	err = json.Unmarshal(body, &currency)
	if err != nil {
		log.Printf("Failed when unmarshal currency: %v\n", err)
		return -1, err
	}

	rows, err := db.Query(`
			SELECT * FROM cards_transaction
			WHERE card_number = $1
			AND EXTRACT(month from date) = $2
			AND EXTRACT(year from date) = $3`,
		cardNumber,
		int(month),
		year,
	)
	if err != nil {
		log.Printf("Failed to get data from database: %v\n", err)
		return -1, err
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
			return -1, err
		}
		cardsTransaction = append(cardsTransaction, cardTransaction)
	}

	var total = 0.0
	for _, transaction := range cardsTransaction {
		if transaction.Currency == `BRL` {
			total += transaction.Amount
		} else {
			total += transaction.Amount * currency.USD
		}
	}

	return total, nil
}
