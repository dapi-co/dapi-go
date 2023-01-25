package response

import (
	"github.com/dapi-co/dapi-go/constants"
)

type IBaseResponse interface {
	GetOperationID() string
	GetSuccess() bool
	GetStatus() constants.ApiStatus
	GetUserInputs() []UserInput
	GetErrType() string
	GetErrMsg() string
	GetCode() int
}

type BaseResponse struct {
	OperationID string              `json:"operationID,omitempty"`
	Success     bool                `json:"success,omitempty"`
	Status      constants.ApiStatus `json:"status,omitempty"`
	UserInputs  []UserInput         `json:"userInputs,omitempty"`
	Type        string              `json:"type,omitempty"`
	Msg         string              `json:"msg,omitempty"`
	Code        int                 `json:"code,omitempty"`
}

func (br *BaseResponse) GetOperationID() string {
	return br.OperationID
}

func (br *BaseResponse) GetSuccess() bool {
	return br.Success
}

func (br *BaseResponse) GetStatus() constants.ApiStatus {
	return br.Status
}

func (br *BaseResponse) GetUserInputs() []UserInput {
	return br.UserInputs
}

func (br *BaseResponse) GetErrType() string {
	return br.Type
}

func (br *BaseResponse) GetErrMsg() string {
	return br.Msg
}

func (br *BaseResponse) GetCode() int {
	return br.Code
}

type ExchangeTokenResponse struct {
	BaseResponse
	AccessToken string `json:"accessToken,omitempty"`
	TokenID     string `json:"tokenID,omitempty"`
	UserID      string `json:"userID,omitempty"`
}

type IdentityResponse struct {
	BaseResponse
	Identity Identity `json:"identity,omitempty"`
}

type AccountsResponse struct {
	BaseResponse
	Accounts []Account `json:"accounts,omitempty"`
}

type BalanceResponse struct {
	BaseResponse
	Balance Balance `json:"balance,omitempty"`
}

type TransactionsResponse struct {
	BaseResponse
	Transactions []Transaction `json:"transactions,omitempty"`
}

type CategorizedTransactionsResponse struct {
	BaseResponse
	Transactions []CategorizedTransaction `json:"transactions,omitempty"`
}

type EnrichedTransactionsResponse struct {
	BaseResponse
	Transactions []EnrichedTransaction `json:"transactions,omitempty"`
}

type BeneficiariesResponse struct {
	BaseResponse
	Beneficiaries []Beneficiary `json:"beneficiaries,omitempty"`
}

type TransferResponse struct {
	BaseResponse
	Reference string `json:"reference,omitempty"`
}

type AccountsMetadataResponse struct {
	BaseResponse
	AccountsMetadata GetAccountsMetadata `json:"accountsMetadata,omitempty"`
}

type GetAchPullResponse struct {
	BaseResponse
	Transfer *GetAchPull `json:"transfer,omitempty"`
}
