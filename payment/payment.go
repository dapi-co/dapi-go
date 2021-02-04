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

// Transfer represents the transfer to be created
type Transfer struct {
	SenderID      string
	ReceiverID    string
	AccountNumber string
	Name          string
	Iban          string
	Amount        float64
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

	body, err := request.DapiRequest(jsonData, constants.GetBeneficiaries, baseHeader)
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
	transfer Transfer,
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

	body, err := request.DapiRequest(jsonData, constants.CreateTransfer, baseHeader)
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
