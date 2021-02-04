package request

import "github.com/dapi-co/dapi-go/response"

// BeneficiaryInfo represents the beneficiary to be created
type BeneficiaryInfo struct {
	Name          string                      `json:"name,omitempty"`
	Type          response.BeneficiaryType    `json:"type,omitempty"`
	Address       response.BeneficiaryAddress `json:"address,omitempty"`
	Country       string                      `json:"country,omitempty"`
	SortCode      string                      `json:"sortCode,omitempty"`
	BranchAddress string                      `json:"branchAddress,omitempty"`
	BankName      string                      `json:"bankName,omitempty"`
	BranchName    string                      `json:"branchName,omitempty"`
	PhoneNumber   string                      `json:"phoneNumber,omitempty"`
	Iban          string                      `json:"iban,omitempty"`
	SwiftCode     string                      `json:"swiftCode,omitempty"`
	AccountNumber string                      `json:"accountNumber,omitempty"`
}
