package handler

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dapi-co/dapi-go/request"
)

var _ http.HandlerFunc = HandleDapiRequests

func HandleDapiRequests(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("an error happened while reading the request body. err: %v\n", err)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(`{"success":false,"msg":"Bad request","type":"BAD_REQUEST","status":"failed"}`))
		return
	}

	resp, err := request.HandleDapiRequest(body, req.Header)
	if err != nil {
		log.Printf("an error happened while handling the request. err: %v\n", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{"success":false,"msg":"Oops! Something happened while performing the request.","type":"UNKNOWN_ERROR","status":"failed"}`))
		return
	}

	rw.Write(resp)
}
