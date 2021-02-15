package auth

import (
	"encoding/json"

	"github.com/dapi-co/dapi-go/config"
	"github.com/dapi-co/dapi-go/constants"
	"github.com/dapi-co/dapi-go/request"
	"github.com/dapi-co/dapi-go/response"
)

// Auth is the base type that allows talking to the auth endpoints
type Auth struct {
	Config *config.Config
}

// ExchangeToken talks to the exchange token endpoint
func (a *Auth) ExchangeToken(
	accessCode string,
	connectionID string) (*response.ExchangeTokenResponse, error) {

	baseRequest := &request.ExchangeTokenRequest{
		AccessCode:   accessCode,
		ConnectionID: connectionID,
		AppSecret:    a.Config.AppSecret,
	}

	jsonData, err := json.Marshal(baseRequest)
	if err != nil {
		return nil, err
	}

	body, err := request.DapiRequest(jsonData, constants.ExchangeToken, nil)
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

// DelinkUser talks to the delink user endpoint
func (a *Auth) DelinkUser(userSecret string) (*response.BaseResponse, error) {

	// config := a.Config

	return nil, nil
}
