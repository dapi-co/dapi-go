package metadata

import (
	"encoding/json"

	"github.com/dapi-co/dapi-go/config"
	"github.com/dapi-co/dapi-go/constants"
	"github.com/dapi-co/dapi-go/request"
	"github.com/dapi-co/dapi-go/response"
)

// Metadata is the base type that allows talking to the metadata endpoints
type Metadata struct {
	Config *config.Config
}

// GetAccountsMetadata talks to the get accounts metadata endpoint
func (m *Metadata) GetAccountsMetadata(
	accessToken string,
	userSecret string,
	userInputs []response.UserInput,
	operationID string,
) (*response.AccountsMetadataResponse, error) {

	baseRequest := &request.BaseRequest{
		UserSecret:  userSecret,
		AppSecret:   m.Config.AppSecret,
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

	body, err := request.DapiRequest(jsonData, constants.GetAccountsMetadata, baseHeader)
	if err != nil {
		return nil, err
	}

	res := response.AccountsMetadataResponse{}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
