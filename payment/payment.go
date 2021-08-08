package payment

import (
	"encoding/json"

	"github.com/dapi-co/dapi-go/config"
	"github.com/dapi-co/dapi-go/constants"
	"github.com/dapi-co/dapi-go/request"
	"github.com/dapi-co/dapi-go/response"
)

// Payment is the base type that allows talking to the payment endpoints
type Payment struct {
	Config *config.Config
}

// CreateTransfer represents the transfer to be created
type CreateTransfer struct {
	Transfer
	ReceiverID    string
	AccountNumber string
	Name          string
	Iban          string
}

// Transfer represents the transfer to be created
type Transfer struct {
	SenderID string
	Amount   float64
	Remark   string
}

// TransferAutoflow represents the transfer to be created
type TransferAutoflow struct {
	Transfer
	Beneficiary request.BeneficiaryInfo
	BankID      string
}

// GetBeneficiaries talks to the get beneficiaries endpoint
func (p *Payment) GetBeneficiaries(
	accessToken string,
	userSecret string,
	userInputs []response.UserInput,
	operationID string,
) (*response.BeneficiariesResponse, error) {

	baseRequest := &request.BaseRequest{
		UserSecret:  userSecret,
		AppSecret:   p.Config.AppSecret,
		UserInputs:  userInputs,
		OperationID: operationID,
	}

	jsonData, err := json.Marshal(baseRequest)
	if err != nil {
		return nil, err
	}

	baseHeader := &request.BaseHeader{
		AccessToken: accessToken,
	}

	body, err := request.DapiRequest(jsonData, constants.DAPI_URL.PAYMENT_URLS.GET_BENEFICIARIES, request.GetHTTPHeader(baseHeader))
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

// CreateTransfer talks to the create transfer endpoint
func (p *Payment) CreateTransfer(
	accessToken string,
	userSecret string,
	transfer CreateTransfer,
	userInputs []response.UserInput,
	operationID string,
) (*response.TransferResponse, error) {

	baseRequest := &request.TransferRequest{
		BaseRequest: request.BaseRequest{
			UserSecret:  userSecret,
			AppSecret:   p.Config.AppSecret,
			UserInputs:  userInputs,
			OperationID: operationID,
		},
		SenderID:      transfer.SenderID,
		ReceiverID:    transfer.ReceiverID,
		Amount:        transfer.Amount,
		Remark:        transfer.Remark,
		Iban:          transfer.Iban,
		AccountNumber: transfer.AccountNumber,
		Name:          transfer.Name,
	}

	jsonData, err := json.Marshal(baseRequest)
	if err != nil {
		return nil, err
	}

	baseHeader := &request.BaseHeader{
		AccessToken: accessToken,
	}

	body, err := request.DapiRequest(jsonData, constants.DAPI_URL.PAYMENT_URLS.CREATE_TRANSFER, request.GetHTTPHeader(baseHeader))
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

// CreateBeneficiary talks to the create beneficiaries endpoint
func (p *Payment) CreateBeneficiary(
	accessToken string,
	userSecret string,
	beneficiary request.CreateBeneficiaryInfo,
	userInputs []response.UserInput,
	operationID string,
) (*response.BaseResponse, error) {

	baseRequest := &request.BeneficiaryRequest{
		BaseRequest: request.BaseRequest{
			UserSecret:  userSecret,
			AppSecret:   p.Config.AppSecret,
			UserInputs:  userInputs,
			OperationID: operationID,
		},
		CreateBeneficiaryInfo: beneficiary,
	}

	jsonData, err := json.Marshal(baseRequest)
	if err != nil {
		return nil, err
	}

	baseHeader := &request.BaseHeader{
		AccessToken: accessToken,
	}

	body, err := request.DapiRequest(jsonData, constants.DAPI_URL.PAYMENT_URLS.CREATE_BENEFICIARY, request.GetHTTPHeader(baseHeader))
	if err != nil {
		return nil, err
	}

	res := response.BaseResponse{}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// TransferAutoflow talks to the transfer autoflow endpoint
func (p *Payment) TransferAutoflow(
	accessToken string,
	userSecret string,
	transfer TransferAutoflow,
	userInputs []response.UserInput,
	operationID string,
) (*response.TransferResponse, error) {

	baseRequest := &request.TransferAutoflowRequest{
		BaseRequest: request.BaseRequest{
			UserSecret:  userSecret,
			AppSecret:   p.Config.AppSecret,
			UserInputs:  userInputs,
			OperationID: operationID,
		},
		SenderID:    transfer.SenderID,
		Amount:      transfer.Amount,
		Remark:      transfer.Remark,
		Beneficiary: transfer.Beneficiary,
		BankID:      transfer.BankID,
	}

	jsonData, err := json.Marshal(baseRequest)
	if err != nil {
		return nil, err
	}

	baseHeader := &request.BaseHeader{
		AccessToken: accessToken,
	}

	body, err := request.DapiRequest(jsonData, constants.DAPI_URL.PAYMENT_URLS.TRANSFER_AUTOFLOW, request.GetHTTPHeader(baseHeader))
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
