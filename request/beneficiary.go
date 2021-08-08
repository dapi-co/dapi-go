package request

import "github.com/dapi-co/dapi-go/response"

// CreateBeneficiaryInfo represents the beneficiary to be created
type CreateBeneficiaryInfo struct {
	BeneficiaryInfo
	Type response.BeneficiaryType `json:"type,omitempty"`
}

// BeneficiaryInfo represents the beneficiary
type BeneficiaryInfo struct {
	Name          string                      `json:"name,omitempty"`
	Nickname      string                      `json:"nickname,omitempty"`
	Address       response.BeneficiaryAddress `json:"address,omitempty"`
	Country       string                      `json:"country,omitempty"`
	BranchAddress string                      `json:"branchAddress,omitempty"`
	BankName      string                      `json:"bankName,omitempty"`
	BranchName    string                      `json:"branchName,omitempty"`
	PhoneNumber   string                      `json:"phoneNumber,omitempty"`
	Iban          string                      `json:"iban,omitempty"`
	RoutingNumber string                      `json:"routingNumber,omitempty"`
	SwiftCode     string                      `json:"swiftCode,omitempty"`
	AccountNumber string                      `json:"accountNumber,omitempty"`
}
