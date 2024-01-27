package restclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Post sends an HTTP POST request to the specified URL with the provided body and headers.
func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	// Marshal the body into JSON format
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	// Create a new HTTP POST request with the URL, JSON body, and provided headers
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = headers

	// Create an HTTP client and send the request
	client := http.Client{}
	return client.Do(request)
}
