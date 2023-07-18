package gateway

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type currencyGatewayHttp struct{}

func NewCurrencyGatewayHttp() *currencyGatewayHttp {
	return &currencyGatewayHttp{}
}

func (c *currencyGatewayHttp) GetCurrencies() (Currency, error) {
	response, err := http.Get(`http://172.17.0.1:3001/currencies`)
	if err != nil {
		log.Printf("Failed when get currency: %v\n", err)
		return Currency{}, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Failed when read response.body: %v\n", err)
		return Currency{}, err
	}

	var currency Currency
	err = json.Unmarshal(body, &currency)
	if err != nil {
		log.Printf("Failed when unmarshal currency: %v\n", err)
		return Currency{}, err
	}

	return currency, nil
}
