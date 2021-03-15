package response

import (
	"github.com/dapi-co/dapi-go/types"
)

type IBaseResponse interface {
	GetStatus() types.ApiStatus
	GetSuccess() bool
	GetOperationID() string
	GetUserInputs() []types.UserInput
	GetErrType() string
	GetErrMsg() string
}

type BaseResponse struct {
	Status      types.ApiStatus   `json:"status,omitempty"`
	Success     bool              `json:"success"`
	OperationID string            `json:"operationID,omitempty"`
	UserInputs  []types.UserInput `json:"userInputs,omitempty"`
	Type        string            `json:"type,omitempty"`
	Msg         string            `json:"msg,omitempty"`
}

func (br *BaseResponse) GetOperationID() string {
	return br.OperationID
}

func (br *BaseResponse) GetSuccess() bool {
	return br.Success
}

func (br *BaseResponse) GetStatus() types.ApiStatus {
	return br.Status
}

func (br *BaseResponse) GetUserInputs() []types.UserInput {
	return br.UserInputs
}

func (br *BaseResponse) GetErrType() string {
	return br.Type
}

func (br *BaseResponse) GetErrMsg() string {
	return br.Msg
}

type ExchangeTokenResponse struct {
	BaseResponse
	AccessToken string `json:"accessToken,omitempty"`
}

type IdentityResponse struct {
	BaseResponse
	Identity types.Identity `json:"identity,omitempty"`
}

type AccountsResponse struct {
	BaseResponse
	Accounts []types.Account `json:"accounts,omitempty"`
}

type BalanceResponse struct {
	BaseResponse
	Balance types.Balance `json:"balance,omitempty"`
}

type TransactionsResponse struct {
	BaseResponse
	Transactions []types.Transaction `json:"transactions,omitempty"`
}

type BeneficiariesResponse struct {
	BaseResponse
	Beneficiaries []types.Beneficiary `json:"beneficiaries,omitempty"`
}

type TransferResponse struct {
	BaseResponse
	Reference string `json:"reference,omitempty"`
}

type AccountsMetadataResponse struct {
	BaseResponse
	AccountsMetadata types.AccountsMetadata `json:"accountsMetadata,omitempty"`
}
