package payment

import (
	"encoding/json"

	"github.com/dapi-co/dapi-go/actions"
	"github.com/dapi-co/dapi-go/config"
	"github.com/dapi-co/dapi-go/request"
	"github.com/dapi-co/dapi-go/response"
	"github.com/dapi-co/dapi-go/types"
)

// Payment is the base type that allows talking to the payment endpoints
type Payment struct {
	Config *config.Config
}

// BeneficiaryInfo represents the beneficiary to be created
type BeneficiaryInfo struct {
	Name          string                   `json:"name,omitempty"`
	Type          types.BeneficiaryType    `json:"type,omitempty"`
	Address       types.BeneficiaryAddress `json:"address,omitempty"`
	Country       string                   `json:"country,omitempty"`
	SortCode      string                   `json:"sortCode,omitempty"`
	BranchAddress string                   `json:"branchAddress,omitempty"`
	BankName      string                   `json:"bankName,omitempty"`
	BranchName    string                   `json:"branchName,omitempty"`
	PhoneNumber   string                   `json:"phoneNumber,omitempty"`
	Iban          string                   `json:"iban,omitempty"`
	SwiftCode     string                   `json:"swiftCode,omitempty"`
	AccountNumber string                   `json:"accountNumber,omitempty"`
}

// Transfer represents the transfer to be created
type Transfer struct {
	SenderID    string
	Amount      float64
	Beneficiary BeneficiaryInfo
}

// GetBeneficiaries talks to the get beneficiaries endpoint
func (p *Payment) GetBeneficiaries(
	tokenID string,
	userID string,
	userSecret string,
	operationID string,
	userInputs []types.UserInput,
) (*response.BeneficiariesResponse, error) {

	req := &request.BaseRequest{
		AppKey:      p.Config.AppKey,
		AppSecret:   p.Config.AppSecret,
		TokenID:     tokenID,
		UserID:      userID,
		UserSecret:  userSecret,
		OperationID: operationID,
		UserInputs:  userInputs,
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	body, err := request.DapiRequest(jsonData, actions.GetBeneficiaries)
	if err != nil {
		return nil, err
	}

	res := response.BeneficiariesResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// transferRequest holds the fields that's needed by the Payment's
// transfer autoflow endpoint.
type transferRequest struct {
	request.BaseRequest
	BundleID    string          `json:"bundleID,omitempty"`
	SenderID    string          `json:"senderID,omitempty"`
	Amount      float64         `json:"amount,omitempty"`
	Beneficiary BeneficiaryInfo `json:"beneficiary,omitempty"`
	HLAPIStep   string          `json:"hlAPIStep,omitempty"`
}

// CreateTransfer talks to the create transfer endpoint
func (p *Payment) CreateTransfer(
	tokenID string,
	userID string,
	userSecret string,
	transfer Transfer,
	hlAPIStep string,
	operationID string,
	userInputs []types.UserInput,
) (*response.TransferResponse, error) {

	baseRequest := &transferRequest{
		BaseRequest: request.BaseRequest{
			AppKey:      p.Config.AppKey,
			AppSecret:   p.Config.AppSecret,
			TokenID:     tokenID,
			UserID:      userID,
			UserSecret:  userSecret,
			OperationID: operationID,
			UserInputs:  userInputs,
		},
		BundleID:    p.Config.BundleID,
		SenderID:    transfer.SenderID,
		Amount:      transfer.Amount,
		Beneficiary: transfer.Beneficiary,
		HLAPIStep:   hlAPIStep,
	}

	jsonData, err := json.Marshal(baseRequest)
	if err != nil {
		return nil, err
	}

	body, err := request.DapiRequest(jsonData, actions.CreateTransfer)
	if err != nil {
		return nil, err
	}

	res := response.TransferResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
