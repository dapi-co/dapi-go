package request

// ACHPullTransfer represents the transfer details
type ACHPullTransfer struct {
	SenderID    string  `json:"senderID"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description,omitempty"`
}
