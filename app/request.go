package app

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/dapi-co/dapi-go/request"
)

func handleDapiRequest(body []byte, header http.Header) ([]byte, error) {
	client := http.Client{}

	// create the request with the provided body
	req, err := http.NewRequest("POST", request.BaseURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// set the header of the request to the provided header, and make sure
	// the content type is set as expected.
	req.Header = header
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
