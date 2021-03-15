package app

type LoginData struct {
	TokenID      string `json:"tokenID,omitempty"`
	UserID       string `json:"userID,omitempty"`
	UserSecret   string `json:"userSecret,omitempty"`
	AccessCode   string `json:"accessCode,omitempty"`
	ConnectionID string `json:"connectionID,omitempty"`
}
