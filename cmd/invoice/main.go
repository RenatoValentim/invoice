package main

import (
	"fmt"
	"invoice/internal/api"
	"invoice/internal/config"
	"invoice/internal/infra/db"
	"invoice/internal/infra/gateway"
	"invoice/internal/usecase"
	"log"

	"github.com/spf13/viper"
)

func main() {
	config.LoadConfig(`.`)

	postgres, err := db.NewPostgresAdapter()
	if err != nil {
		log.Printf("Failed to connect the database: %v\n", err)
	}

	transactionDAODatabase := db.NewTransactionDAODatabase(postgres)

	currencyBaseUrl := fmt.Sprintf(
		`%s:%d`,
		viper.GetString(`currency_host`),
		viper.GetInt(`currency_port`),
	)
	currecyGatewayHttp := gateway.NewCurrencyGatewayHttp(currencyBaseUrl)
	calculateInvoice := usecase.NewCalculateInvoice(
		transactionDAODatabase,
		currecyGatewayHttp,
	)
	api.InvoiceController(*calculateInvoice)
}
