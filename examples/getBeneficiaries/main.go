package main

import (
	"log"

	"github.com/dapi-co/dapi-go/app"
	"github.com/dapi-co/dapi-go/config"
	"github.com/dapi-co/dapi-go/types"
)

func main() {
	// Populate the Config fields with your AppKey and your AppSecret.
	// Note: The AppSecret is required only if your app handles the AppSecret
	//       by itself.
	myConfig := config.Config{
		AppKey:    "",
		AppSecret: "",
	}

	// Populate the LoginData fields with the corresponding values got from logging in
	myLoginData := app.LoginData{
		TokenID:      "",
		UserID:       "",
		UserSecret:   "",
		AccessCode:   "",
		ConnectionID: "",
	}

	// Create the app with the corresponding config and login data
	myApp := app.NewDapiApp(myConfig, myLoginData)

	// Populate these two fields only if you are continuing with a previous call
	// to submit userInputs.
	operationID := ""
	userInputs := []types.UserInput{}

	// Call the endpoint and print the got response or error
	resp, err := myApp.GetBeneficiaries(operationID, userInputs)
	if err != nil {
		log.Fatalf("GetBeneficiaries error: %#v", err)
	}
	log.Printf("GetBeneficiaries response: %#v", resp)
}
