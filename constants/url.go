package constants

// TODO: change this to the actual DD URL
const BaseURL = "http://localhost:8081"

type DapiAction string

const (
	ExchangeToken       DapiAction = "/auth/ExchangeToken"
	GetIdentity         DapiAction = "/data/identity/get"
	GetAccounts         DapiAction = "/data/accounts/get"
	GetBalance          DapiAction = "/data/balance/get"
	GetTransactions     DapiAction = "/data/transactions/get"
	CreateTransfer      DapiAction = "/payment/transfer/autoflow"
	GetBeneficiaries    DapiAction = "/payment/beneficiaries/get"
	GetAccountsMetadata DapiAction = "/metadata/accounts/get"
	DelinkUser          DapiAction = "/users/delinkuser"
	OperationStatus     DapiAction = "/operation/status"
)
