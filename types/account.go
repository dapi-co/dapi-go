package types

type AccountType string

const (
	CurrentAccount AccountType = "current"
	SavingsAccount AccountType = "savings"
	LoanAccount    AccountType = "loan"
	CreditAccount  AccountType = "credit"
	DepositAccount AccountType = "deposit"
	OtherAccount   AccountType = "other"
)

type Account struct {
	Iban        string   `json:"iban,omitempty"`
	Number      string   `json:"number,omitempty"`
	Currency    Currency `json:"currency,omitempty"`
	Type        string   `json:"type,omitempty"`
	ID          string   `json:"id,omitempty"`
	IsFavourite bool     `json:"isFavourite,omitempty"`
	Name        string   `json:"name,omitempty"`
	Balance     Balance  `json:"balance,omitempty"`
}
