package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type transactionDAODatabase struct{}

func NewTransactionDAODatabase() *transactionDAODatabase {
	return &transactionDAODatabase{}
}

func (t *transactionDAODatabase) GetTransactions(cardNumber string, month, year int) ([]CardTransaction, error) {
	dsn := fmt.Sprintf(
		`host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo`,
		viper.GetString(`db_host`),
		viper.GetString(`db_user`),
		viper.GetString(`db_password`),
		viper.GetString(`db_name`),
		viper.GetString(`db_port`),
	)

	db, err := sql.Open(`postgres`, dsn)
	if err != nil {
		log.Printf("Failed to connect on database: %v\n", err)
		return nil, err
	}
	defer db.Close()

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
		return nil, err
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
			return nil, err
		}
		cardsTransaction = append(cardsTransaction, cardTransaction)
	}

	return cardsTransaction, nil
}
