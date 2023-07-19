package usecase_test

import (
	"invoice/internal/config"
	db_contracts "invoice/internal/infra/db/contracts"
	gateway_contracts "invoice/internal/infra/gateway/contracts"
	"invoice/internal/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TransactionDAOFake struct{}

func (t TransactionDAOFake) GetTransactions(cardNumber string, month, year int) ([]db_contracts.CardTransaction, error) {
	return []db_contracts.CardTransaction{
		{
			Amount:   100,
			Currency: "BRL",
		},
		{
			Amount:   1000,
			Currency: "BRL",
		},
		{
			Amount:   600,
			Currency: "USD",
		},
	}, nil
}

type CurrencyGatewayFake struct{}

func (c CurrencyGatewayFake) GetCurrencies() (gateway_contracts.Currency, error) {
	return gateway_contracts.Currency{
		USD: 2.0,
	}, nil
}

func TestCalculateInvoice(t *testing.T) {
	assert := assert.New(t)
	config.LoadConfig(`../../`)

	t.Run(`Should to calculate invoice`, func(t *testing.T) {
		var transactionDAOFake TransactionDAOFake
		var currencyGatewayFake CurrencyGatewayFake

		calculateInvoice := usecase.NewCalculateInvoice(
			transactionDAOFake,
			currencyGatewayFake,
		)
		total, _ := calculateInvoice.Execute(`1234`)

		assert.Equal(2300.0, total)
	})
}
