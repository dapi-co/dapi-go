package response

type CategorizedTransaction struct {
	Amount       float64         `json:"amount,omitempty"`
	Date         string          `json:"date,omitempty"`
	Type         TransactionType `json:"type,omitempty"`
	Description  string          `json:"description,omitempty"`
	Details      string          `json:"details,omitempty"`
	Currency     Currency        `json:"currency,omitempty"`
	BeforeAmount float64         `json:"beforeAmount,omitempty"`
	AfterAmount  float64         `json:"afterAmount,omitempty"`
	Reference    string          `json:"reference,omitempty"`
	Category     string          `json:"category,omitempty"`
}
