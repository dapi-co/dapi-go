package types

type Balance struct {
	Amount        float64  `json:"amount,omitempty"`
	Currency      Currency `json:"currency,omitempty"`
	AccountNumber string   `json:"accountNumber,omitempty"`
}
