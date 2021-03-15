package types

type PhoneNumberType string

const (
	Mobile PhoneNumberType = "mobile"
	Home   PhoneNumberType = "home"
	Office PhoneNumberType = "office"
	Fax    PhoneNumberType = "fax"
)

type PhoneNumber struct {
	Type  PhoneNumberType `json:"type,omitempty"`
	Value string          `json:"value,omitempty"`
}
