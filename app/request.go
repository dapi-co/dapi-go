package app

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/dapi-co/dapi-go/constants"
)

func handleDapiRequest(body []byte, header http.Header) ([]byte, error) {
	client := http.Client{}

	// create the request with the provided body
	request, err := http.NewRequest("POST", constants.BaseURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// set the header of the request to the provided header, and make sure
	// the content type is set as expected.
	request.Header = header
	request.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(request)
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
