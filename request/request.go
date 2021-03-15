package request

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dapi-co/dapi-go/actions"
	"github.com/dapi-co/dapi-go/types"
)

const BaseURL = "https://dd.dapi.co"

// BaseRequest holds the fields that's needed by all endpoints
type BaseRequest struct {
	AppKey      string            `json:"appKey,omitempty"`
	AppSecret   string            `json:"appSecret,omitempty"`
	TokenID     string            `json:"tokenID,omitempty"`
	UserID      string            `json:"userID,omitempty"`
	UserSecret  string            `json:"userSecret,omitempty"`
	OperationID string            `json:"operationID,omitempty"`
	UserInputs  []types.UserInput `json:"userInputs,omitempty"`
}

// DapiRequest creates a request to the API, on the product specified by productURL,
// with the body of the request set as the provided body, and the headers as the
// provided headers.
func DapiRequest(body []byte, action actions.DapiAction) ([]byte, error) {
	client := http.Client{}

	// unmarshal the body to a map, to add the action to it
	bodyMap := make(map[string]interface{})
	err := json.Unmarshal(body, &bodyMap)
	if err != nil {
		return nil, err
	}
	bodyMap["action"] = action

	// marshal the new body back to json
	body, err = json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", BaseURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
