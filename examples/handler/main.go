package main

import (
	"log"
	"net/http"

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

	// Start the app server to handle requests specific to this app.
	// You can now make request to the specific port, by only passing
	// the 'action' field in the request body.
	err := http.ListenAndServe(":8085", http.HandlerFunc(myApp.HandleDapiRequests))
	if err != nil {
		log.Fatal(err)
	}
}
