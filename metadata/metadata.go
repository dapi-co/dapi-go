package metadata

import (
	"encoding/json"

	"github.com/dapi-co/dapi-go/actions"
	"github.com/dapi-co/dapi-go/config"
	"github.com/dapi-co/dapi-go/request"
	"github.com/dapi-co/dapi-go/response"
	"github.com/dapi-co/dapi-go/types"
)

// Metadata is the base type that allows talking to the metadata endpoints
type Metadata struct {
	Config *config.Config
}

// GetAccountsMetadata talks to the get accounts metadata endpoint
func (m *Metadata) GetAccountsMetadata(
	tokenID string,
	userSecret string,
	operationID string,
	userInputs []types.UserInput,
) (*response.AccountsMetadataResponse, error) {

	req := &request.BaseRequest{
		AppKey:      m.Config.AppKey,
		AppSecret:   m.Config.AppSecret,
		TokenID:     tokenID,
		UserID:      "",
		UserSecret:  userSecret,
		OperationID: operationID,
		UserInputs:  userInputs,
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	body, err := request.DapiRequest(jsonData, actions.GetAccountsMetadata)
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
