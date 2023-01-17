package request

// AchPullTransferInfo represents the transfer details
type AchPullTransferInfo struct {
	SenderID    string  `json:"senderID"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description,omitempty"`
}
