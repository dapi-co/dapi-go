package response

type BeneficiaryType string

const (
	Same          BeneficiaryType = "same"
	Local         BeneficiaryType = "local"
	International BeneficiaryType = "intl"
)

type BeneficiaryStatus string

const (
	Approved  BeneficiaryStatus = "approved"
	Rejected  BeneficiaryStatus = "rejected"
	Cancelled BeneficiaryStatus = "cancelled"
	Pending   BeneficiaryStatus = "waiting_for_confirmation"
	Modified  BeneficiaryStatus = "modified_for_pending_approval"
)

type BeneficiaryAddress struct {
	Line1 string `json:"line1,omitempty"`
	Line2 string `json:"line2,omitempty"`
	Line3 string `json:"line3,omitempty"`
}

type Beneficiary struct {
	Iban          string            `json:"iban,omitempty"`
	AccountNumber string            `json:"accountNumber,omitempty"`
	Status        BeneficiaryStatus `json:"status,omitempty"`
	Type          BeneficiaryType   `json:"type,omitempty"`
	ID            string            `json:"id,omitempty"`
	Name          string            `json:"name,omitempty"`
}
