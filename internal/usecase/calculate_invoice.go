package usecase

import (
	"invoice/internal/infra/db/contracts"
	"invoice/internal/infra/gateway"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type CalculateInvoice struct {
	transactionDAO  contracts.TransactionDAO
	currencyGateway gateway.CurrencyGateway
}

func NewCalculateInvoice(
	transactionDAO contracts.TransactionDAO,
	currencyGateway gateway.CurrencyGateway,
) *CalculateInvoice {
	return &CalculateInvoice{
		transactionDAO:  transactionDAO,
		currencyGateway: currencyGateway,
	}
}

func (c *CalculateInvoice) Execute(cardNumber string) (float64, error) {
	currentDate := time.Now()
	month := int(currentDate.Month())
	year := currentDate.Year()

	currency, err := c.currencyGateway.GetCurrencies()
	if err != nil {
		log.Printf("Failed when get currency: %v\n", err)
		return -1, err
	}

	cardsTransaction, err := c.transactionDAO.GetTransactions(
		cardNumber,
		month,
		year,
	)
	if err != nil {
		log.Printf("Failed when searh cards transaction: %v\n", err)
		return -1, err
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
