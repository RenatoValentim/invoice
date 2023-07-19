package db

import (
	"invoice/internal/infra/db/contracts"
	"log"
)

type transactionDAODatabase struct {
	conn contracts.Connection
}

func NewTransactionDAODatabase(conn contracts.Connection) *transactionDAODatabase {
	return &transactionDAODatabase{
		conn: conn,
	}
}

func (t *transactionDAODatabase) GetTransactions(cardNumber string, month, year int) ([]contracts.CardTransaction, error) {
	rows, err := t.conn.Query(`
			SELECT * FROM cards_transaction
			WHERE card_number = $1
			AND EXTRACT(month from date) = $2
			AND EXTRACT(year from date) = $3`,
		cardNumber,
		int(month),
		year,
	)
	defer t.conn.Close()
	if err != nil {
		log.Printf("Failed to get data from database: %v\n", err)
		return nil, err
	}

	var cardsTransaction []contracts.CardTransaction
	for rows.Next() {
		var cardTransaction contracts.CardTransaction
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
