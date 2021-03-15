package data

import (
	"encoding/json"
	"time"

	"github.com/dapi-co/dapi-go/actions"
	"github.com/dapi-co/dapi-go/config"
	"github.com/dapi-co/dapi-go/request"
	"github.com/dapi-co/dapi-go/response"
	"github.com/dapi-co/dapi-go/types"
)

// Data is the base type that allows talking to the data endpoints
type Data struct {
	Config *config.Config
}

// GetIdentity talks to the get identity endpoint
func (d *Data) GetIdentity(
	tokenID string,
	userID string,
	userSecret string,
	operationID string,
	userInputs []types.UserInput,
) (*response.IdentityResponse, error) {

	req := &request.BaseRequest{
		AppKey:      d.Config.AppKey,
		AppSecret:   d.Config.AppSecret,
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

	body, err := request.DapiRequest(jsonData, actions.GetIdentity)
	if err != nil {
		return nil, err
	}

	res := response.IdentityResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetAccounts talks to the get accounts endpoint
func (d *Data) GetAccounts(
	tokenID string,
	userID string,
	userSecret string,
	operationID string,
	userInputs []types.UserInput,
) (*response.AccountsResponse, error) {

	req := &request.BaseRequest{
		AppKey:      d.Config.AppKey,
		AppSecret:   d.Config.AppSecret,
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

	body, err := request.DapiRequest(jsonData, actions.GetAccounts)
	if err != nil {
		return nil, err
	}

	res := response.AccountsResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// BalanceRequest holds the fields that's needed by the Data's
// get balance endpoint.
type balanceRequest struct {
	request.BaseRequest
	AccountID string `json:"accountID"`
}

// GetBalance talks to the get balance endpoint
func (d *Data) GetBalance(
	tokenID string,
	userID string,
	userSecret string,
	accountID string,
	operationID string,
	userInputs []types.UserInput,
) (*response.BalanceResponse, error) {

	req := &balanceRequest{
		BaseRequest: request.BaseRequest{
			AppKey:      d.Config.AppKey,
			AppSecret:   d.Config.AppSecret,
			TokenID:     tokenID,
			UserID:      userID,
			UserSecret:  userSecret,
			OperationID: operationID,
			UserInputs:  userInputs,
		},
		AccountID: accountID,
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	body, err := request.DapiRequest(jsonData, actions.GetBalance)
	if err != nil {
		return nil, err
	}

	res := response.BalanceResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// transactionsRequest holds the fields that's needed by the Data's
// get transactions endpoint.
type transactionsRequest struct {
	request.BaseRequest
	AccountID string `json:"accountID"`
	FromDate  string `json:"fromDate"`
	ToDate    string `json:"toDate"`
}

// GetTransactions talks to the get transactions endpoint
func (d *Data) GetTransactions(
	tokenID string,
	userID string,
	userSecret string,
	accountID string,
	fromDate time.Time,
	toDate time.Time,
	operationID string,
	userInputs []types.UserInput,
) (*response.TransactionsResponse, error) {

	dateFormat := "2006-01-02"

	req := &transactionsRequest{
		BaseRequest: request.BaseRequest{
			AppKey:      d.Config.AppKey,
			AppSecret:   d.Config.AppSecret,
			TokenID:     tokenID,
			UserID:      userID,
			UserSecret:  userSecret,
			OperationID: operationID,
			UserInputs:  userInputs,
		},
		AccountID: accountID,
		FromDate:  fromDate.Format(dateFormat),
		ToDate:    toDate.Format(dateFormat),
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	body, err := request.DapiRequest(jsonData, actions.GetTransactions)
	if err != nil {
		return nil, err
	}

	res := response.TransactionsResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
