package main

import (
	"log"
	"net/http"

	"github.com/dapi-co/dapi-go/app"
	"github.com/dapi-co/dapi-go/config"
)

func main() {

	myApp := &app.DapiApp{
		Config: &config.Config{
			AppSecret: "MY_APP_SECRET",
		},
	}
	err := http.ListenAndServe(":8085", http.HandlerFunc(myApp.HandleSDKDapiRequests))
	if err != nil {
		log.Fatal(err)
	}
}

// set the appSecret for your app in the AppSecret field of the myConfig var
// above, then run the main function above, then call: http://localhost:8085/
// and set the 'action' field in the body to the endpoint you want to call,
// and add the values that's wanted by that endpoint.
