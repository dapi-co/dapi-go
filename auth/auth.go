package auth

import (
	"encoding/json"

	"github.com/dapi-co/dapi-go/actions"
	"github.com/dapi-co/dapi-go/config"
	"github.com/dapi-co/dapi-go/request"
	"github.com/dapi-co/dapi-go/response"
)

// Auth is the base type that allows talking to the auth endpoints
type Auth struct {
	Config *config.Config
}

// exchangeTokenRequest holds the fields that's needed by the Auth's
// exchange token endpoint.
type exchangeTokenRequest struct {
	AppKey       string `json:"appKey"`
	AppSecret    string `json:"appSecret"`
	TokenID      string `json:"tokenID"`
	AccessCode   string `json:"accessCode"`
	ConnectionID string `json:"connectionID"`
}

// ExchangeToken talks to the exchange token endpoint
func (a *Auth) ExchangeToken(
	tokenID string,
	accessCode string,
	connectionID string,
) (*response.ExchangeTokenResponse, error) {

	req := &exchangeTokenRequest{
		AppKey:       a.Config.AppKey,
		AppSecret:    a.Config.AppSecret,
		TokenID:      tokenID,
		AccessCode:   accessCode,
		ConnectionID: connectionID,
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	body, err := request.DapiRequest(jsonData, actions.ExchangeToken)
	if err != nil {
		return nil, err
	}

	res := response.ExchangeTokenResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
