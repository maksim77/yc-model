package ycmodel

import "encoding/json"

type Response struct {
	// HTTP status code
	StatusCode int `json:"statusCode"`

	// Map containins HTTP request headers and their values
	// If the same header is passed multiple times, the map contains the last value passed.
	Headers map[string]string `json:"headers"`

	// Map containins HTTP request headers and list of their values.
	// It contains the same keys as Headers, but if any header was repeated several times,
	// the list for it will contain all the passed values for this header.
	MultiValueHeaders map[string][]string `json:"multiValueHeaders"`

	// Body of request in string format.
	// The data can be encoded in Base64 format (in this case IsBase64Encoded wil be set to true)
	Body            string `json:"body"`
	IsBase64Encoded bool   `json:"isBase64Encoded"`
}

func (r *Response) Marshal() ([]byte, error) {
	return json.Marshal(&r)
}
