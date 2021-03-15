package types

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/dapi-co/dapi-go/utils"
)

type Country struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

type TransferBounds struct {
	Minimum  int             `json:"minimum,omitempty"`
	Currency Currency        `json:"currency,omitempty"`
	Type     BeneficiaryType `json:"type,omitempty"`
}

type ValidatorProps struct {
	Required          bool        `json:"required,omitempty"`
	Optional          bool        `json:"optional,omitempty"`
	Length            int         `json:"length,omitempty"`
	AllowedCharacters string      `json:"allowedCharacters,omitempty"`
	Attributes        interface{} `json:"attributes,omitempty"`
}

type AddressValidatorProps struct {
	Length int            `json:"length,omitempty"`
	Line1  ValidatorProps `json:"line1,omitempty"`
	Line2  ValidatorProps `json:"line2,omitempty"`
	Line3  ValidatorProps `json:"line3,omitempty"`
}

type CreateBeneficiaryValidatorProps struct {
	Name          ValidatorProps        `json:"name,omitempty"`
	Nickname      ValidatorProps        `json:"nickname,omitempty"`
	SwiftCode     ValidatorProps        `json:"swiftCode,omitempty"`
	Iban          ValidatorProps        `json:"iban,omitempty"`
	AccountNumber ValidatorProps        `json:"accountNumber,omitempty"`
	Address       AddressValidatorProps `json:"address,omitempty"`
	BranchAddress ValidatorProps        `json:"branchAddress,omitempty"`
	BranchName    ValidatorProps        `json:"branchName,omitempty"`
	Country       ValidatorProps        `json:"country,omitempty"`
	PhoneNumber   ValidatorProps        `json:"phoneNumber,omitempty"`
	SortCode      ValidatorProps        `json:"sortCode,omitempty"`
}

type CreateTransferValidatorProps struct {
	Remarks ValidatorProps `json:"remarks,omitempty"`
}

type CreateBeneficiaryValidator struct {
	Local CreateBeneficiaryValidatorProps `json:"local,omitempty"`
	Same  CreateBeneficiaryValidatorProps `json:"same,omitempty"`
}

type CreateTransferValidator struct {
	Local CreateTransferValidatorProps `json:"local,omitempty"`
	Same  CreateTransferValidatorProps `json:"same,omitempty"`
}

type Validators struct {
	CreateBeneficiary CreateBeneficiaryValidator `json:"createBeneficiary,omitempty"`
	CreateTransfer    CreateTransferValidator    `json:"createTransfer,omitempty"`
}

type Range struct {
	Value int    `json:"value,omitempty"`
	Unit  string `json:"unit,omitempty"`
}

type AccountsMetadata struct {
	SwiftCode                                          string             `json:"swiftCode,omitempty"`
	SortCode                                           string             `json:"sortCode,omitempty"`
	BankName                                           string             `json:"bankName,omitempty"`
	BranchName                                         string             `json:"branchName,omitempty"`
	BranchAddress                                      string             `json:"branchAddress,omitempty"`
	Address                                            BeneficiaryAddress `json:"address,omitempty"`
	TransferBounds                                     []TransferBounds   `json:"transferBounds,omitempty"`
	BeneficiaryCoolDownPeriod                          Range              `json:"beneficiaryCoolDownPeriod,omitempty"`
	TransactionRange                                   Range              `json:"transactionRange,omitempty"`
	Country                                            Country            `json:"country,omitempty"`
	IsCreateBeneficiaryEndpointRequired                bool               `json:"isCreateBeneficiaryEndpointRequired,omitempty"`
	WillNewlyAddedBeneficiaryExistBeforeCoolDownPeriod bool               `json:"willNewlyAddedBeneficiaryExistBeforeCoolDownPeriod,omitempty"`
	Validators                                         Validators         `json:"validators,omitempty"`
}

//IsValid returns an error if the property is invalid according to the validator settings
func (v *ValidatorProps) IsValid(propName, propValue string) error {

	if !v.Required && propValue == "" {
		return nil
	}

	if v.Required && propValue == "" {
		return errors.New(`Property '` + propName + `' can not be empty`)
	}

	if len(propValue) > v.Length {
		return errors.New(`Property '` + propName + `' length can not exceed ` + fmt.Sprint(v.Length))
	}

	if len(v.AllowedCharacters) > 0 {
		if isValid, _ := regexp.MatchString("^("+v.AllowedCharacters+")*$", propValue); !isValid {
			return errors.New(`Property '` + propName + `' contains invalid characters. Allowed characters are: ` + v.AllowedCharacters)
		}
	}

	return nil
}

//TruncateIfNeeded truncates the value if longer than the max length
//and randomizes the last 5 characters
func (v *ValidatorProps) TruncateIfNeeded(propValue string) string {

	//TODO: This regex is WRONG. It hardcodes the allowedCharacters string, while the one in the validator
	//object should be used -_-
	regex, _ := regexp.Compile("[^a-zA-Z0-9 ]")
	//regex, _ := regexp.Compile("^(" + v.AllowedCharacters + ")*$")

	propValue = regex.ReplaceAllString(propValue, "")
	if v.Length == 0 || len(propValue) <= v.Length {
		return propValue
	}

	return propValue[:v.Length-5] + utils.GenerateRandomString(5)
}
