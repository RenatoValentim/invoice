package gateway

import (
	"encoding/json"
	"fmt"
	"invoice/internal/infra/gateway/contracts"
	"io/ioutil"
	"log"
	"net/http"
)

type currencyGatewayHttp struct {
	BaseUrl string
}

func NewCurrencyGatewayHttp(baseUrl string) *currencyGatewayHttp {
	return &currencyGatewayHttp{
		BaseUrl: baseUrl,
	}
}

func (c *currencyGatewayHttp) GetCurrencies() (contracts.Currency, error) {
	url := fmt.Sprintf(`%s/currencies`, c.BaseUrl)
	response, err := http.Get(url)
	if err != nil {
		log.Printf("Failed when get currency: %v\n", err)
		return contracts.Currency{}, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Failed when read response.body: %v\n", err)
		return contracts.Currency{}, err
	}

	var currency contracts.Currency
	err = json.Unmarshal(body, &currency)
	if err != nil {
		log.Printf("Failed when unmarshal currency: %v\n", err)
		return contracts.Currency{}, err
	}

	return currency, nil
}
