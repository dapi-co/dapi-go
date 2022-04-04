# dapi-go

A client library that talks to the [Dapi](https://dapi.com) [API](https://api.dapi.com).

## Quickstart

### Configure Project

First add the library module to your project.

```
go get github.com/dapi-co/dapi-go
```

### Configure Library

1. Create a Dapi app and products instances with your App Secret.

```go
package main

import (
	"github.com/dapi-co/dapi-go/app"
	"github.com/dapi-co/dapi-go/auth"
	"github.com/dapi-co/dapi-go/config"
	"github.com/dapi-co/dapi-go/data"
	"github.com/dapi-co/dapi-go/metadata"
	"github.com/dapi-co/dapi-go/payment"
)

func main() {
	// create a config object that holds the secret of this app
	myAppConfig := &config.Config{
		AppSecret: "YOUR_APP_SECRET",
	}

	// init a DapiApp instance
	myApp := &app.DapiApp{
		Config: myAppConfig,
	}

	// init the products you want to use
	myAuth := auth.Auth{Config: myAppConfig}
	myData := data.Data{Config: myAppConfig}
	myPayment := payment.Payment{Config: myAppConfig}
	myMetadata := metadata.Metadata{Config: myAppConfig}

	// use any of the myApp, myAuth, myData, myPayment, or myMetadata methods..
}
```

2. Now you can use any of the functions of any product (`Data` for example) instance, `myData`, to call Dapi with your `appSecret`.

```go
package main

import (
	"github.com/dapi-co/dapi-go/app"
	"github.com/dapi-co/dapi-go/config"
	"github.com/dapi-co/dapi-go/data"
)

func main() {
	// create a config object that holds the secret of this app
	myAppConfig := &config.Config{
		AppSecret: "YOUR_APP_SECRET",
	}

	// init the product you want to use
	myData := data.Data{Config: myAppConfig}

	// use any of the `myData` methods..

	// provide the operationID and the userInputs only if needed
	accountsResp, err := myData.GetAccounts("YOUR_ACCESS_TOKEN", "YOUR_USER_SECRET", nil, "")
	if err != nil {
		// handle the error..
		return
	}
	if accountsResp.Status != constants.StatusDone {
		// handle the unsuccessful response..
		return
	}

	accounts := accountsResp.Accounts
	// use the got accounts array..
}
```

3. Or, you can use the `DapiApp` instance to handle requests to a specific endpoint in your server. Our code will basically update the request to add your app's `appSecret` to it, and forward the request to Dapi, then respond back with the got response.

```go
package main

import (
	"log"
	"net/http"

	"github.com/dapi-co/dapi-go/app"
	"github.com/dapi-co/dapi-go/config"
)

func main() {
	// create a config object that holds the secret of this app
	myAppConfig := &config.Config{
		AppSecret: "YOUR_APP_SECRET",
	}

	// init the a DapiApp instance
	myApp := &app.DapiApp{
		Config: myAppConfig,
	}

	// start a simple server that listens on the provided endpoint, and handles requests
	// through the handler of the DapiApp instance.
	err := http.ListenAndServe("YOUR_ADDRESS", http.HandlerFunc(myApp.HandleSDKDapiRequests))
	if err != nil {
		log.Fatal(err)
	}
}
```

## Reference

### BaseResponse

All the responses extend BaseResponse class. Meaning all the responses described below in the document will have following fields besides the ones specific to each response

| Parameter | Type | Description |
|---|---|---|
| OperationID | `string` | Unique ID generated to identify a specific operation. |
| Success | `bool` | Returns true if request is successful and false otherwise." |
| Status | `ApiStatus` (Enum) | The status of the job. <br><br> `done` - Operation Completed. <br> `failed` - Operation Failed. <br> `user_input_required` - Pending User Input. <br> `initialized` - Operation In Progress. <br><br> For further explanation see [Operation Statuses](https://dapi-api.readme.io/docs/operation-statuses). |
| UserInputs | `[]UserInput` | Specifies the type of further information required from the user before the job can be completed. <br><br> Note: It's only returned if operation status is `user_input_required` |
| Type | `string` | Type of error encountered. <br><br> Note: It's only returned if operation status is `failed` |
| Msg | `string` | Detailed description of the error. <br><br> Note: It's only returned if operation status is `failed` |

#### UserInput Object

| Parameter | Type | Description |
|---|---|---|
| Id | `UserInputID` (Enum) | Type of input required. <br><br> You can read more about user input types on [User Input Types](https://dapi-api.readme.io/docs/user-input-types). |
| Query | `string` | Textual description of what is required from the user side. |
| Index | `int` | Is used in case more than one user input is requested. <br> Will always be 0 If only one input is requested. |
| Answer | `string` | User input that must be submitted. In the response it will always be empty. |

### Methods

#### Auth.ExchangeToken

Method is used to obtain user's permanent access token by exchanging it with access code received during the user authentication (user login).

##### Note:

You can read more about how to obtain a permanent token on [Obtain an Access Token](https://dapi-api.readme.io/docs/get-an-access-token).

##### Method Description

```go
func (*Auth) ExchangeToken(accessCode string, connectionID string) (*response.ExchangeTokenResponse, error)
```

##### Input Parameters

| Parameter | Type | Description |
|---|---|---|
| **accessCode** <br> _REQUIRED_ | `string` | Unique code for a user’s successful login to **Connect**. Returned in the response of **UserLogin**. |
| **connectionID** <br> _REQUIRED_ | `string` | The `connectionID` from a user’s successful log in to **Connect**. |

##### Response

In addition to the fields described in the BaseResponse, it has the following fields, which will only be returned if the status is `done`:

| Parameter | Type | Description |
|---|---|---|
| **AccessToken** | `string` | A unique permanent token linked to the user. |

---

#### Data.GetIdentity

Method is used to retrieve personal details about the user.

##### Method Description

```go
func (*Data) GetIdentity(accessToken string, userSecret string, userInputs []response.UserInput, operationID string) (*response.IdentityResponse, error)
```

##### Input Parameters

| Parameter | Type | Description |
|---|---|---|
| **accessToken** <br> _REQUIRED_ | `string` | Access Token obtained using the `ExchangeToken` method. |
| **userSecret** <br> _REQUIRED_ | `string` | The `userSecret` from a user’s successful log in to **Connect**. |
| **operationID** <br> _OPTIONAL_ | `string` | The `OperationID` from a previous call's response. <br> Required only when resuming a previous call that responded with `user_input_required` status, to provided user inputs. |
| **userInputs** <br> _OPTIONAL_ | `[]UserInput` | Array of `UserInput` object, that are needed to complete this operation. <br> Required only if a previous call responded with `user_input_required` status. <br><br> You can read more about user inputs specification on [Specify User Input](https://dapi-api.readme.io/docs/specify-user-input) |

###### UserInput Object

| Parameter | Type | Description |
|---|---|---|
| Id | `UserInputID` (Enum) | Type of input required. <br><br> You can read more about user input types on [User Input Types](https://dapi-api.readme.io/docs/user-input-types). |
| Index | `int` | Is used in case more than one user input is requested. <br> Will always be 0 If only one input is requested. |
| Answer | `string` | User input that must be submitted. |

##### Response

In addition to the fields described in the BaseResponse, it has the following fields, which will only be returned if the status is `done`:

| Parameter | Type | Description |
|---|---|---|
| Identity | `Identity` | An object containing the identity data of the user. |

---

#### Data.GetAccounts

Method is used to retrieve list of all the bank accounts registered on the user. The list will contain all types of bank accounts.

##### Method Description

```go
func (*Data) GetAccounts(accessToken string, userSecret string, userInputs []response.UserInput, operationID string) (*response.AccountsResponse, error)
```

##### Input Parameters

| Parameter | Type | Description |
|---|---|---|
| **accessToken** <br> _REQUIRED_ | `string` | Access Token obtained using the `ExchangeToken` method. |
| **userSecret** <br> _REQUIRED_ | `string` | The `userSecret` from a user’s successful log in to **Connect**. |
| **operationID** <br> _OPTIONAL_ | `string` | The `OperationID` from a previous call's response. <br> Required only when resuming a previous call that responded with `user_input_required` status, to provided user inputs. |
| **userInputs** <br> _OPTIONAL_ | `[]UserInput` | Array of `UserInput` object, that are needed to complete this operation. <br> Required only if a previous call responded with `user_input_required` status. <br><br> You can read more about user inputs specification on [Specify User Input](https://dapi-api.readme.io/docs/specify-user-input) |

###### UserInput Object

| Parameter | Type | Description |
|---|---|---|
| id | `UserInputID` (Enum) | Type of input required. <br><br> You can read more about user input types on [User Input Types](https://dapi-api.readme.io/docs/user-input-types). |
| index | `int` | Is used in case more than one user input is requested. <br> Will always be 0 If only one input is requested. |
| answer | `string` | User input that must be submitted. |

##### Response

In addition to the fields described in the BaseResponse, it has the following fields, which will only be returned if the status is `done`:

| Parameter | Type | Description |
|---|---|---|
| Accounts | `[]Account` | An array containing the accounts data of the user. |

---

#### Data.GetBalance

Method is used to retrieve balance on specific bank account of the user.

##### Method Description

```go
func (*Data) GetBalance(accessToken string, userSecret string, accountID string, userInputs []response.UserInput, operationID string) (*response.BalanceResponse, error)
```

##### Input Parameters

| Parameter | Type | Description |
|---|---|---|
| **accountID** <br> _REQUIRED_ | `string` | The bank account ID which its balance is requested. <br> Retrieved from one of the accounts returned from the `GetAccounts` method. |
| **accessToken** <br> _REQUIRED_ | `string` | Access Token obtained using the `ExchangeToken` method. |
| **userSecret** <br> _REQUIRED_ | `string` | The `userSecret` from a user’s successful log in to **Connect**. |
| **operationID** <br> _OPTIONAL_ | `string` | The `OperationID` from a previous call's response. <br> Required only when resuming a previous call that responded with `user_input_required` status, to provided user inputs. |
| **userInputs** <br> _OPTIONAL_ | `[]UserInput` | Array of `UserInput` object, that are needed to complete this operation. <br> Required only if a previous call responded with `user_input_required` status. <br><br> You can read more about user inputs specification on [Specify User Input](https://dapi-api.readme.io/docs/specify-user-input) |

###### UserInput Object

| Parameter | Type | Description |
|---|---|---|
| id | `UserInputID` (Enum) | Type of input required. <br><br> You can read more about user input types on [User Input Types](https://dapi-api.readme.io/docs/user-input-types). |
| index | `int` | Is used in case more than one user input is requested. <br> Will always be 0 If only one input is requested. |
| answer | `string` | User input that must be submitted. |

##### Response

In addition to the fields described in the BaseResponse, it has the following fields, which will only be valid if the status is `done`:

| Parameter | Type | Description |
|---|---|---|
| Balance | `Balance` | An object containing the account's balance information. |

---

#### Data.GetTransactions

Method is used to retrieve transactions that user has performed over a specific period of time from their bank account. The transaction list is unfiltered, meaning the response will contain all the transactions performed by the user (not just the transactions performed using your app).

Date range of the transactions that can be retrieved varies for each bank. The range supported by the users bank is shown in the response parameter `transactionRange` of Get Accounts Metadata endpoint.

##### Method Description

```go
func (*Data) GetTransactions(accessToken string, userSecret string, accountID string, fromDate time.Time, toDate time.Time, userInputs []response.UserInput, operationID string) (*response.TransactionsResponse, error)
```

##### Input Parameters

| Parameter | Type | Description |
|---|---|---|
| **accountID** <br> _REQUIRED_ | `string` | The bank account ID which its transactions are requested. <br> Retrieved from one of the accounts returned from the `getAccounts` method. |
| **fromDate** <br> _REQUIRED_ | `time.Time` | The start date of the transactions wanted. |
| **toDate** <br> _REQUIRED_ | `time.Time` | The end date of the transactions wanted. |
| **accessToken** <br> _REQUIRED_ | `string` | Access Token obtained using the `ExchangeToken` method. |
| **userSecret** <br> _REQUIRED_ | `string` | The `userSecret` from a user’s successful log in to **Connect**. |
| **operationID** <br> _OPTIONAL_ | `string` | The `OperationID` from a previous call's response. <br> Required only when resuming a previous call that responded with `user_input_required` status, to provided user inputs. |
| **userInputs** <br> _OPTIONAL_ | `[]UserInput` | Array of `UserInput` object, that are needed to complete this operation. <br> Required only if a previous call responded with `user_input_required` status. <br><br> You can read more about user inputs specification on [Specify User Input](https://dapi-api.readme.io/docs/specify-user-input) |

###### UserInput Object

| Parameter | Type | Description |
|---|---|---|
| id | `UserInputID` (Enum) | Type of input required. <br><br> You can read more about user input types on [User Input Types](https://dapi-api.readme.io/docs/user-input-types). |
| index | `int` | Is used in case more than one user input is requested. <br> Will always be 0 If only one input is requested. |
| answer | `string` | User input that must be submitted. |

##### Response

In addition to the fields described in the BaseResponse, it has the following fields, which will only be valid if the status is `done`:

| Parameter | Type | Description |
|---|---|---|
| Transactions | `[]Transaction` | Array containing the transactional data for the specified account within the specified period. |

---

#### Payment.GetBeneficiaries

Method is used to retrieve list of all the beneficiaries already added for a user within a financial institution.

##### Method Description

```go
func (*Payment) GetBeneficiaries(accessToken string, userSecret string, userInputs []response.UserInput, operationID string) (*response.BeneficiariesResponse, error)
```

##### Input Parameters

| Parameter | Type | Description |
|---|---|---|
| **accessToken** <br> _REQUIRED_ | `string` | Access Token obtained using the `ExchangeToken` method. |
| **userSecret** <br> _REQUIRED_ | `string` | The `userSecret` from a user’s successful log in to **Connect**. |
| **operationID** <br> _OPTIONAL_ | `string` | The `OperationID` from a previous call's response. <br> Required only when resuming a previous call that responded with `user_input_required` status, to provided user inputs. |
| **userInputs** <br> _OPTIONAL_ | `[]UserInput` | Array of `UserInput` object, that are needed to complete this operation. <br> Required only if a previous call responded with `user_input_required` status. <br><br> You can read more about user inputs specification on [Specify User Input](https://dapi-api.readme.io/docs/specify-user-input) |

###### UserInput Object

| Parameter | Type | Description |
|---|---|---|
| id | `UserInputID` (Enum) | Type of input required. <br><br> You can read more about user input types on [User Input Types](https://dapi-api.readme.io/docs/user-input-types). |
| index | `int` | Is used in case more than one user input is requested. <br> Will always be 0 If only one input is requested. |
| answer | `string` | User input that must be submitted. |

##### Response

In addition to the fields described in the BaseResponse, it has the following fields, which will only be returned if the status is `done`:

| Parameter | Type | Description |
|---|---|---|
| Beneficiaries | `[]Beneficiary` | An array containing the beneficiary information. |

---

#### Payment.CreateBeneficiary

Method is used to retrieve list of all the beneficiaries already added for a user within a financial institution.

##### Method Description

```go
func (*Payment) CreateBeneficiary(accessToken string, userSecret string, beneficiary request.CreateBeneficiaryInfo, userInputs []response.UserInput, operationID string) (*response.BaseResponse, error)
```

##### Input Parameters

| Parameter | Type | Description |
|---|---|---|
| **beneficiary** <br> _REQUIRED_ | `request.CreateBeneficiaryInfo` | An object that contains info about the beneficiary that should be added. |
| **accessToken** <br> _REQUIRED_ | `string` | Access Token obtained using the `ExchangeToken` method. |
| **userSecret** <br> _REQUIRED_ | `string` | The `userSecret` from a user’s successful log in to **Connect**. |
| **operationID** <br> _OPTIONAL_ | `string` | The `OperationID` from a previous call's response. <br> Required only when resuming a previous call that responded with `user_input_required` status, to provided user inputs. |
| **userInputs** <br> _OPTIONAL_ | `[]UserInput` | Array of `UserInput` object, that are needed to complete this operation. <br> Required only if a previous call responded with `user_input_required` status. <br><br> You can read more about user inputs specification on [Specify User Input](https://dapi-api.readme.io/docs/specify-user-input) |

###### UserInput Object

| Parameter | Type | Description |
|---|---|---|
| id | `UserInputID` (Enum) | Type of input required. <br><br> You can read more about user input types on [User Input Types](https://dapi-api.readme.io/docs/user-input-types). |
| index | `int` | Is used in case more than one user input is requested. <br> Will always be 0 If only one input is requested. |
| answer | `string` | User input that must be submitted. |

###### request.CreateBeneficiaryInfo Object

| Parameter | Type | Description |
|---|---|---|
| **Name** <br> _REQUIRED_ | `string` | Name of the beneficiary. |
| **AccountNumber** <br> _REQUIRED_ | `string` | Account number of the beneficiary. |
| **Iban** <br> _REQUIRED_ | `string` | Beneficiary's IBAN number. |
| **SwiftCode** <br> _REQUIRED_ | `string` | Beneficiary's financial institution's SWIFT code. |
| **Type** <br> _REQUIRED_ | `response.BeneficiaryType` (Enum) | Type of beneficiary. <br> For further explanation see [Beneficiary Types](https://dapi-api.readme.io/docs/beneficiary-types). |
| **Address** <br> _REQUIRED_ | `response.BeneficiaryAddress` | An object containing the address information of the beneficiary. |
| **Country** <br> _REQUIRED_ | `string` | Name of the country in all uppercase letters. |
| **BranchAddress** <br> _REQUIRED_ | `string` | Address of the financial institution’s specific branch. |
| **BranchName** <br> _REQUIRED_ | `string` | Name of the financial institution’s specific branch. |
| **PhoneNumber** <br> _OPTIONAL_ | `string` | Beneficiary's phone number. |
| **RoutingNumber** <br> _OPTIONAL_ | `string` | Beneficiary's Routing number, needed only for US banks accounts. |

###### response.BeneficiaryAddress Object

| Parameter | Type | Description |
|---|---|---|
| **Line1** <br> _REQUIRED_ | `string` | Street name and number. Note: value should not contain any commas or special characters. |
| **Line2** <br> _REQUIRED_ | `string` | City name. Note: value should not contain any commas or special characters. |
| **Line3** <br> _REQUIRED_ | `string` | Country name. Note: value should not contain any commas or special characters. |

##### Response

Method returns only the fields defined in the BaseResponse.

---

#### Payment.CreateTransfer

Method is used to initiate a new payment from one account to another account.

##### Important

We suggest you use `TransferAutoflow` method instead to initiate a payment. `TransferAutoFlow` abstracts all the validations and processing logic, required to initiate a transaction using `CreateTransfer` method.

You can read about `TransferAutoFlow` further in the document.

##### Method Description

```go
func (*Payment) CreateTransfer(accessToken string, userSecret string, transfer CreateTransfer, userInputs []response.UserInput, operationID string) (*response.TransferResponse, error)
```

##### Input Parameters

| Parameter | Type | Description |
|---|---|---|
| **transfer** <br> _REQUIRED_ | `CreateTransfer` | An object that contains info about the transfer that should be initiated. |
| **accessToken** <br> _REQUIRED_ | `string` | Access Token obtained using the `ExchangeToken` method. |
| **userSecret** <br> _REQUIRED_ | `string` | The `userSecret` from a user’s successful log in to **Connect**. |
| **operationID** <br> _OPTIONAL_ | `string` | The `OperationID` from a previous call's response. <br> Required only when resuming a previous call that responded with `user_input_required` status, to provided user inputs. |
| **userInputs** <br> _OPTIONAL_ | `[]UserInput` | Array of `UserInput` object, that are needed to complete this operation. <br> Required only if a previous call responded with `user_input_required` status. <br><br> You can read more about user inputs specification on [Specify User Input](https://dapi-api.readme.io/docs/specify-user-input) |

###### UserInput Object

| Parameter | Type | Description |
|---|---|---|
| id | `UserInputID` (Enum) | Type of input required. <br><br> You can read more about user input types on [User Input Types](https://dapi-api.readme.io/docs/user-input-types). |
| index | `int` | Is used in case more than one user input is requested. <br> Will always be 0 If only one input is requested. |
| answer | `string` | User input that must be submitted. |

###### CreateTransfer Object

| Parameter | Type | Description |
|---|---|---|
| **SenderID** <br> _REQUIRED_ | `string` | The id of the account which the money should be sent from. <br> Retrieved from one of the accounts array returned from the getAccounts method. |
| **Amount** <br> _REQUIRED_ | `float64` | The amount of money which should be sent. |
| **ReceiverID** <br> _OPTIONAL_ | `string` | The id of the beneficiary which the money should be sent to. <br> Retrieved from one of the beneficiaries array returned from the getBeneficiaries method. <br> Needed only when creating a transfer from a bank that requires the receiver to be already registered as a beneficiary to perform a transaction. |
| **Name** <br> _OPTIONAL_ | `string` | The name of receiver. <br> Needed only when creating a transfer from a bank that handles the creation of beneficiaries on its own, internally, and doesn't require the receiver to be already registered as a beneficiary to perform a transaction. |
| **AccountNumber** <br> _OPTIONAL_ | `string` | The Account Number of the receiver's account. <br> Needed only when creating a transfer from a bank that handles the creation of beneficiaries on its own, internally, and doesn't require the receiver to be already registered as a beneficiary to perform a transaction. |
| **Iban** <br> _OPTIONAL_ | `string` | The IBAN of the receiver's account. <br> Needed only when creating a transfer from a bank that handles the creation of beneficiaries on its own, internally, and doesn't require the receiver to be already registered as a beneficiary to perform a transaction. |

##### Response

In addition to the fields described in the BaseResponse, it has the following fields, which will only be returned if the status is `done`:

| Parameter | Type | Description |
|---|---|---|
| reference | `String` | Transaction reference string returned by the bank. |

---

#### Payment.TransferAutoflow

Method is used to initiate a new payment from one account to another account, without having to care nor handle any special cases or scenarios.

##### Method Description

```go
func (*Payment) TransferAutoflow(accessToken string, userSecret string, transfer TransferAutoflow, userInputs []response.UserInput, operationID string) (*response.TransferResponse, error)
```

##### Input Parameters

| Parameter | Type | Description |
|---|---|---|
| **transfer** <br> _REQUIRED_ | `TransferAutoflow` | An object that contains info about the transfer that should be initiated, and any other details that's used to automate the operation. |
| **accessToken** <br> _REQUIRED_ | `string` | Access Token obtained using the `ExchangeToken` method. |
| **userSecret** <br> _REQUIRED_ | `string` | The `userSecret` from a user’s successful log in to **Connect**. |
| **operationID** <br> _OPTIONAL_ | `string` | The `OperationID` from a previous call's response. <br> Required only when resuming a previous call that responded with `user_input_required` status, to provided user inputs. |
| **userInputs** <br> _OPTIONAL_ | `[]UserInput` | Array of `UserInput` object, that are needed to complete this operation. <br> Required only if a previous call responded with `user_input_required` status. <br><br> You can read more about user inputs specification on [Specify User Input](https://dapi-api.readme.io/docs/specify-user-input) |

###### UserInput Object

| Parameter | Type | Description |
|---|---|---|
| id | `UserInputID` (Enum) | Type of input required. <br><br> You can read more about user input types on [User Input Types](https://dapi-api.readme.io/docs/user-input-types). |
| index | `int` | Is used in case more than one user input is requested. <br> Will always be 0 If only one input is requested. |
| answer | `string` | User input that must be submitted. |

###### TransferAutoflow Object

| Parameter | Type | Description |
|---|---|---|
| **SenderID** <br> _REQUIRED_ | `string` | The id of the account which the money should be sent from. <br> Retrieved from one of the accounts array returned from the getAccounts method. |
| **Amount** <br> _REQUIRED_ | `float64` | The amount of money which should be sent. |
| **Beneficiary** <br> _REQUIRED_ | `request.BeneficiaryInfo` | An object that holds the info about the beneficiary which the money should be sent to. |
| **BankID** <br> _REQUIRED_ | `string` | The bankID of the user which is initiating this transfer. |

###### request.BeneficiaryInfo Object

| Parameter | Type | Description |
|---|---|---|
| **Name** <br> _REQUIRED_ | `string` | Name of the beneficiary. |
| **AccountNumber** <br> _REQUIRED_ | `string` | Account number of the beneficiary. |
| **Iban** <br> _REQUIRED_ | `string` | Beneficiary's IBAN number. |
| **SwiftCode** <br> _REQUIRED_ | `string` | Beneficiary's financial institution's SWIFT code. |
| **Address** <br> _REQUIRED_ | `response.BeneficiaryAddress` | An object containing the address information of the beneficiary. |
| **Country** <br> _REQUIRED_ | `string` | Name of the country in all uppercase letters. |
| **BranchAddress** <br> _REQUIRED_ | `string` | Address of the financial institution’s specific branch. |
| **BranchName** <br> _REQUIRED_ | `string` | Name of the financial institution’s specific branch. |
| **PhoneNumber** <br> _OPTIONAL_ | `string` | Beneficiary's phone number. |
| **RoutingNumber** <br> _OPTIONAL_ | `string` | Beneficiary's Routing number, needed only for US banks accounts. |

##### Response

In addition to the fields described in the BaseResponse, it has the following fields, which will only be returned if the status is `done`:

| Parameter | Type | Description |
|---|---|---|
| reference | `String` | Transaction reference string returned by the bank. |

---
