package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dapi-co/dapi-go/auth"
	"github.com/dapi-co/dapi-go/config"
	"github.com/dapi-co/dapi-go/data"
	"github.com/dapi-co/dapi-go/payment"
	"github.com/dapi-co/dapi-go/response"
	"github.com/dapi-co/dapi-go/types"
)

type DapiApp struct {
	config    config.Config
	loginData LoginData
}

func NewDapiApp(config config.Config, loginData LoginData) *DapiApp {
	return &DapiApp{config: config, loginData: loginData}
}

func (app *DapiApp) ExchangeToken() (*response.ExchangeTokenResponse, error) {
	a := auth.Auth{Config: &app.config}
	return a.ExchangeToken(app.loginData.TokenID, app.loginData.AccessCode, app.loginData.ConnectionID)
}

func (app *DapiApp) GetIdentity(
	operationID string,
	userInputs []types.UserInput,
) (*response.IdentityResponse, error) {
	d := data.Data{Config: &app.config}
	return d.GetIdentity(
		app.loginData.TokenID, app.loginData.UserID, app.loginData.UserSecret,
		operationID, userInputs,
	)
}

func (app *DapiApp) GetAccounts(
	operationID string,
	userInputs []types.UserInput,
) (*response.AccountsResponse, error) {
	d := data.Data{Config: &app.config}
	return d.GetAccounts(
		app.loginData.TokenID, app.loginData.UserID, app.loginData.UserSecret,
		operationID, userInputs,
	)
}

func (app *DapiApp) GetBalance(
	accountID string,
	operationID string,
	userInputs []types.UserInput,
) (*response.BalanceResponse, error) {
	d := data.Data{Config: &app.config}
	return d.GetBalance(
		app.loginData.TokenID, app.loginData.UserID, app.loginData.UserSecret,
		accountID, operationID, userInputs,
	)
}

func (app *DapiApp) GetTransactions(
	accountID string,
	fromDate time.Time,
	toDate time.Time,
	operationID string,
	userInputs []types.UserInput,
) (*response.TransactionsResponse, error) {
	d := data.Data{Config: &app.config}
	return d.GetTransactions(
		app.loginData.TokenID, app.loginData.UserID, app.loginData.UserSecret,
		accountID, fromDate, toDate, operationID, userInputs,
	)
}

func (app *DapiApp) GetBeneficiaries(
	operationID string,
	userInputs []types.UserInput,
) (*response.BeneficiariesResponse, error) {
	p := payment.Payment{Config: &app.config}
	return p.GetBeneficiaries(
		app.loginData.TokenID, app.loginData.UserID, app.loginData.UserSecret,
		operationID, userInputs,
	)
}

func (app *DapiApp) CreateTransfer(
	transfer payment.Transfer,
	hlAPIStep string,
	operationID string,
	userInputs []types.UserInput,
) (*response.TransferResponse, error) {
	p := payment.Payment{Config: &app.config}
	return p.CreateTransfer(
		app.loginData.TokenID, app.loginData.UserID, app.loginData.UserSecret,
		transfer, hlAPIStep, operationID, userInputs,
	)
}

// HandleDapiRequests is an HTTP handler function, which redirects all requests
// to Dapi's API, after adding the fields specific to this App to their body.
// The only required field in the received request is the 'action' field.
func (app *DapiApp) HandleDapiRequests(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	// read the body sent
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(`{"success":false,"msg":"Bad request","type":"BAD_REQUEST","status":"failed"}`))
		return
	}

	// unmarshal the body to a map
	bodyMap := make(map[string]interface{})
	err = json.Unmarshal(body, &bodyMap)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(`{"success":false,"msg":"Bad request","type":"BAD_REQUEST","status":"failed"}`))
		return
	}

	// add the fields specific to this app to the body
	bodyMap["appKey"] = app.config.AppKey
	bodyMap["appSecret"] = app.config.AppSecret
	bodyMap["tokenID"] = app.loginData.TokenID
	bodyMap["userID"] = app.loginData.UserID
	bodyMap["userSecret"] = app.loginData.UserSecret
	bodyMap["accessCode"] = app.loginData.AccessCode
	bodyMap["connectionID"] = app.loginData.ConnectionID

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
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{"success":false,"msg":"Oops! Something happened while performing the request.","type":"UNKNOWN_ERROR","status":"failed"}`))
		return
	}

	rw.Write(resp)
}
