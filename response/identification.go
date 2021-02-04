package response

type IDType string

const (
	Passport   IDType = "passport"
	NationalID IDType = "national_id"
	VisaNumber IDType = "visa_number"
)

type Identification struct {
	Type  IDType `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}
