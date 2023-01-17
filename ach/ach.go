package ach

import (
	"encoding/json"

	"github.com/dapi-co/dapi-go/config"
	"github.com/dapi-co/dapi-go/constants"
	"github.com/dapi-co/dapi-go/request"
	"github.com/dapi-co/dapi-go/response"
)

// Payment is the base type that allows talking to the payment endpoints
type Ach struct {
	Config *config.Config
}

// CreatePull represents the pull transfer to be created
type CreatePull struct {
	Transfer request.AchPullTransferInfo
}

// CreatePull talks to the create pull endpoint
func (a *Ach) CreatePull(
	accessToken string,
	userSecret string,
	pullTransfer CreatePull,
	userInputs []response.UserInput,
	operationID string,
) (*response.BaseResponse, error) {

	baseRequest := &request.AchPullRequest{
		BaseRequest: request.BaseRequest{
			UserSecret:  userSecret,
			AppSecret:   a.Config.AppSecret,
			UserInputs:  userInputs,
			OperationID: operationID,
		},
		Transfer: pullTransfer.Transfer,
	}

	jsonData, err := json.Marshal(baseRequest)
	if err != nil {
		return nil, err
	}

	baseHeader := &request.BaseHeader{
		AccessToken: accessToken,
	}

	body, err := request.DapiRequest(jsonData, constants.DAPI_URL.ACH_URLS.CREATE_PULL, request.GetHTTPHeader(baseHeader))
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
