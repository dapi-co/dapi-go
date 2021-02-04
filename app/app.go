package app

import (
	"time"

	"github.com/dapi-co/dapi-go/auth"
	"github.com/dapi-co/dapi-go/config"
	"github.com/dapi-co/dapi-go/data"
	"github.com/dapi-co/dapi-go/payment"
	"github.com/dapi-co/dapi-go/response"
)

type DapiApp struct {
	config config.Config
}

func NewDapiApp(config config.Config) *DapiApp {
	return &DapiApp{config: config}
}

func (app *DapiApp) ExchangeToken(
	accessCode string,
	connectionID string,
) (*response.ExchangeTokenResponse, error) {
	a := auth.Auth{Config: &app.config}
	return a.ExchangeToken(accessCode, connectionID)
}

// TODO: add DelinkUser

func (app *DapiApp) GetIdentity(
	accessToken string,
	userSecret string,
	userInputs []response.UserInput,
	operationID string,
) (*response.IdentityResponse, error) {
	d := data.Data{Config: &app.config}
	return d.GetIdentity(accessToken, userSecret, userInputs, operationID)
}

func (app *DapiApp) GetAccounts(
	accessToken string,
	userSecret string,
	userInputs []response.UserInput,
	operationID string,
) (*response.AccountsResponse, error) {
	d := data.Data{Config: &app.config}
	return d.GetAccounts(accessToken, userSecret, userInputs, operationID)
}

func (app *DapiApp) GetBalance(
	accessToken string,
	userSecret string,
	accountID string,
	userInputs []response.UserInput,
	operationID string,
) (*response.BalanceResponse, error) {
	d := data.Data{Config: &app.config}
	return d.GetBalance(accessToken, userSecret, accountID, userInputs, operationID)
}

func (app *DapiApp) GetTransactions(
	accessToken string,
	userSecret string,
	accountID string,
	fromDate time.Time,
	toDate time.Time,
	userInputs []response.UserInput,
	operationID string,
) (*response.TransactionsResponse, error) {
	d := data.Data{Config: &app.config}
	return d.GetTransactions(accessToken, userSecret, accountID, fromDate, toDate, userInputs, operationID)
}

func (app *DapiApp) GetBeneficiaries(
	accessToken string,
	userSecret string,
	userInputs []response.UserInput,
	operationID string,
) (*response.BeneficiariesResponse, error) {
	p := payment.Payment{Config: &app.config}
	return p.GetBeneficiaries(accessToken, userSecret, userInputs, operationID)
}

func (app *DapiApp) CreateTransfer(
	accessToken string,
	userSecret string,
	transfer payment.Transfer,
	userInputs []response.UserInput,
	operationID string,
) (*response.TransferResponse, error) {
	p := payment.Payment{Config: &app.config}
	return p.CreateTransfer(accessToken, userSecret, transfer, userInputs, operationID)
}
