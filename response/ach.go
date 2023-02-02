package response

type ACHPull struct {
	Amount   float64  `json:"amount,omitempty"`
	Currency Currency `json:"currency,omitempty"`
	Status   string   `json:"status,omitempty"`
}
