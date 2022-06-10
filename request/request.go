package request

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dapi-co/dapi-go/constants"
	"github.com/dapi-co/dapi-go/response"
)

// BaseRequest holds the fields that's needed by all endpoints
type BaseRequest struct {
	UserSecret  string               `json:"userSecret"`
	AppSecret   string               `json:"appSecret"`
	UserInputs  []response.UserInput `json:"userInputs,omitempty"`
	OperationID string               `json:"operationID,omitempty"`
}

// ExchangeTokenRequest holds the fields that's needed by the Auth's
// exchange token endpoint.
type ExchangeTokenRequest struct {
	AppSecret    string `json:"appSecret"`
	AccessCode   string `json:"accessCode"`
	ConnectionID string `json:"connectionID"`
}

// BalanceRequest holds the fields that's needed by the Data's
// get balance endpoint.
type BalanceRequest struct {
	BaseRequest
	AccountID string `json:"accountID"`
}

// TransactionsRequest holds the fields that's needed by the Data's
// get transactions endpoint.
type TransactionsRequest struct {
	BaseRequest
	AccountID string `json:"accountID"`
	FromDate  string `json:"fromDate"`
	ToDate    string `json:"toDate"`
}

// TransferRequest holds the fields that's needed by the Payment's
// create transfer endpoint.
type TransferRequest struct {
	BaseRequest
	SenderID      string  `json:"senderID"`
	ReceiverID    string  `json:"receiverID,omitempty"`
	Amount        float64 `json:"amount"`
	Remark        string  `json:"remark,omitempty"`
	Iban          string  `json:"iban,omitempty"`
	AccountNumber string  `json:"accountNumber,omitempty"`
	Name          string  `json:"name,omitempty"`
}

// TransferAutoflowRequest holds the fields that's needed by the Payment's
// transfer autoflow endpoint.
type TransferAutoflowRequest struct {
	BaseRequest
	SenderID    string  `json:"senderID"`
	Amount      float64 `json:"amount"`
	Remark      string  `json:"remark,omitempty"`
	Beneficiary BeneficiaryInfo
	BankID      string `json:"bankID"`
}

// BeneficiaryRequest holds the fields that's needed by the Payment's
// create beneficiaries endpoint.
type BeneficiaryRequest struct {
	BaseRequest
	CreateBeneficiaryInfo
}

type NoHeader struct{}

// BaseHeader holds any fields that's needed in the header of the request
type BaseHeader struct {
	AccessToken string
}

// DapiRequest creates a request to the API, on the product specified by productURL,
// with the body of the request set as the provided body, and the headers as the
// provided headers.
func DapiRequest(body []byte, action string, header http.Header) ([]byte, error) {
	client := http.Client{Timeout: 5 * time.Minute}

	url := constants.DAPI_URL.BASE_URL + action

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if header != nil {
		request.Header = header
	}

	request.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(request)
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

// DapiSDKRequest forwards the passed body and header to Dapi's DD url, and
// return the response body and nil, if no errors happened, otherwise it returns
// nil and the error happened.
func DapiSDKRequest(body []byte, header http.Header) ([]byte, error) {
	client := http.Client{Timeout: 5 * time.Minute}

	request, err := http.NewRequest("POST", constants.DD_URL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if header != nil {
		request.Header = header
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Host", "dd.dapi.com")

	resp, err := client.Do(request)
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

func GetHTTPHeader(header *BaseHeader) http.Header {
	httpHeader := http.Header{}
	authHeader := ""
	if header.AccessToken != "" {
		authHeader = "Bearer " + header.AccessToken
	}
	httpHeader.Add("Authorization", authHeader)

	return httpHeader
}
