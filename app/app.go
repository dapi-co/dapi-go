package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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

func (app *DapiApp) HandleDapiRequests(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	// read the body sent
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(`{"success":false,"msg":"Bad request","type":"BAD_REQUEST","status":"failed"}`))
		return
	}

	// unmarshal the body to a map, to add the appSecret to it
	bodyMap := make(map[string]interface{})
	err = json.Unmarshal(body, &bodyMap)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(`{"success":false,"msg":"Bad request","type":"BAD_REQUEST","status":"failed"}`))
		return
	}
	bodyMap["appSecret"] = app.config.AppSecret

	// marshal the new body back to json
	body, err = json.Marshal(bodyMap)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{"success":false,"msg":"Oops! Something happened while performing the request.","type":"UNKNOWN_ERROR","status":"failed"}`))
		return
	}

	// forward the request to be handled
	resp, err := handleDapiRequest(body, req.Header)
	if err != nil {
		log.Printf("an error happened while handling the request. err: %v\n", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{"success":false,"msg":"Oops! Something happened while performing the request.","type":"UNKNOWN_ERROR","status":"failed"}`))
		return
	}

	rw.Write(resp)
}
