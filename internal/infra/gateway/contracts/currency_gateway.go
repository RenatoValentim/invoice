package contracts

type Currency struct {
	USD float64 `json:"usd"`
}

type CurrencyGateway interface {
	GetCurrencies() (Currency, error)
}
