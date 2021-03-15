package main

import (
	"log"

	"github.com/dapi-co/dapi-go/app"
	"github.com/dapi-co/dapi-go/config"
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

	// Call the endpoint and print the got response or error
	resp, err := myApp.ExchangeToken()
	if err != nil {
		log.Fatalf("ExchangeToken error: %#v", err)
	}
	log.Printf("ExchangeToken response: %#v", resp)
}
