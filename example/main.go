package main

import (
	"log"
	"net/http"

	"github.com/dapi-co/dapi-go/handler"
)

func main() {

	err := http.ListenAndServe(":8085", http.HandlerFunc(handler.HandleDapiRequests))
	if err != nil {
		log.Fatal(err)
	}
}

// run the main function above, then call: http://localhost:8085/
// and set the 'action' field in the body to the endpoint you want to call,
// and add the values that's wanted by that endpoint.
