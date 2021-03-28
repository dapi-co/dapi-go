package data

import (
	"encoding/json"
	"time"

	"github.com/dapi-co/dapi-go/config"
	"github.com/dapi-co/dapi-go/constants"
	"github.com/dapi-co/dapi-go/request"
	"github.com/dapi-co/dapi-go/response"
)

// Data is the base type that allows talking to the data endpoints
type Data struct {
	Config *config.Config
}

// GetIdentity talks to the get identity endpoint
func (d *Data) GetIdentity(
	accessToken string,
	userSecret string,
	userInputs []response.UserInput,
	operationID string,
) (*response.IdentityResponse, error) {

	baseRequest := &request.BaseRequest{
		UserSecret:  userSecret,
		AppSecret:   d.Config.AppSecret,
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

	body, err := request.DapiRequest(jsonData, constants.DAPI_URL.DATA_URLS.GET_IDENTITY, request.GetHTTPHeader(baseHeader))

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
	accessToken string,
	userSecret string,
	userInputs []response.UserInput,
	operationID string,
) (*response.AccountsResponse, error) {

	baseRequest := &request.BaseRequest{
		UserSecret:  userSecret,
		AppSecret:   d.Config.AppSecret,
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

	body, err := request.DapiRequest(jsonData, constants.DAPI_URL.DATA_URLS.GET_ACCOUNTS, request.GetHTTPHeader(baseHeader))
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

// GetBalance talks to the get balance endpoint
func (d *Data) GetBalance(
	accessToken string,
	userSecret string,
	accountID string,
	userInputs []response.UserInput,
	operationID string,
) (*response.BalanceResponse, error) {

	balanceRequest := &request.BalanceRequest{
		BaseRequest: request.BaseRequest{
			UserSecret:  userSecret,
			AppSecret:   d.Config.AppSecret,
			UserInputs:  userInputs,
			OperationID: operationID,
		},
		AccountID: accountID,
	}

	jsonData, err := json.Marshal(balanceRequest)
	if err != nil {
		return nil, err
	}

	baseHeader := &request.BaseHeader{
		AccessToken: accessToken,
	}

	body, err := request.DapiRequest(jsonData, constants.DAPI_URL.DATA_URLS.GET_BALANCE, request.GetHTTPHeader(baseHeader))
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

// GetTransactions talks to the get transactions endpoint
func (d *Data) GetTransactions(
	accessToken string,
	userSecret string,
	accountID string,
	fromDate time.Time,
	toDate time.Time,
	userInputs []response.UserInput,
	operationID string,
) (*response.TransactionsResponse, error) {

	dateFormat := "2006-01-02"

	transactionsRequest := &request.TransactionsRequest{
		BaseRequest: request.BaseRequest{
			UserSecret:  userSecret,
			AppSecret:   d.Config.AppSecret,
			UserInputs:  userInputs,
			OperationID: operationID,
		},
		AccountID: accountID,
		FromDate:  fromDate.Format(dateFormat),
		ToDate:    toDate.Format(dateFormat),
	}

	jsonData, err := json.Marshal(transactionsRequest)
	if err != nil {
		return nil, err
	}

	baseHeader := &request.BaseHeader{
		AccessToken: accessToken,
	}

	body, err := request.DapiRequest(jsonData, constants.DAPI_URL.DATA_URLS.GET_TRANSACTIONS, request.GetHTTPHeader(baseHeader))
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
