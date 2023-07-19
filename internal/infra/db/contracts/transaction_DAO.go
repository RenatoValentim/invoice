package contracts

type CardTransaction struct {
	CardNumber  string
	Description string
	Amount      float64
	Currency    string
	Date        string
}

type TransactionDAO interface {
	GetTransactions(cardNumber string, month, year int) ([]CardTransaction, error)
}
