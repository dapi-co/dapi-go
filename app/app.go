package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dapi-co/dapi-go/config"
	"github.com/dapi-co/dapi-go/request"
)

type DapiApp struct {
	Config *config.Config
}

func (app *DapiApp) HandleSDKDapiRequests(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	// read the body sent
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(`{"success":false,"msg":"Bad request","type":"BAD_REQUEST","status":"failed"}`))
		return
	}

	// unmarshal the body to a map, to add the appSecret to it
	bodyMap := make(map[string]interface{})
	err = json.Unmarshal(body, &bodyMap)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(`{"success":false,"msg":"Bad request","type":"BAD_REQUEST","status":"failed"}`))
		return
	}
	bodyMap["appSecret"] = app.Config.AppSecret

	// marshal the new body back to json
	body, err = json.Marshal(bodyMap)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{"success":false,"msg":"Oops! Something happened while performing the request.","type":"UNKNOWN_ERROR","status":"failed"}`))
		return
	}

	// forward the request to be handled
	resp, err := request.DapiSDKRequest(body, req.Header)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{"success":false,"msg":"Oops! Something happened while performing the request.","type":"UNKNOWN_ERROR","status":"failed"}`))
		return
	}

	rw.Write(resp)
}
