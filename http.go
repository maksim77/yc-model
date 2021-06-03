package ycmodel

import "encoding/base64"

// Identity is user identification
type Identity struct {
	SourceIP  string `json:"sourceIp"`
	UserAgent string `json:"userAgent"`
}

type RequestContext struct {
	Identity Identity `json:"identity"`

	// HTTP method (DELETE, GET, HEAD, OPTIONS, PATCH, POST or PUT)
	HTTPMethod string `json:"httpMethod"`

	// Request ID generated in the router
	RequestId string `json:"requestId"`

	// Request time in CLF format
	RequestTime string `json:"requestTime"`

	// Request Time in unixtime format
	RequestTimeEpoch int64 `json:"requestTimeEpoch"`
}

type Request struct {
	// HTTP method (DELETE, GET, HEAD, OPTIONS, PATCH, POST or PUT)
	HTTPMethod string `json:"httpMethod"`

	// Map containins HTTP request headers and their values
	// If the same header is passed multiple times, the map contains the last value passed.
	Headers map[string]string `json:"headers"`

	// Map containins HTTP request headers and list of their values.
	// It contains the same keys as Headers, but if any header was repeated several times,
	// the list for it will contain all the passed values for this header.
	MultiValueHeaders map[string][]string `json:"multiValueHeaders"`

	// Map contains query string parameters and their values.
	// If the same parameter is passed multiple times, the map contains the last value passed.
	QueryStringParameters map[string]string `json:"queryStringParameters"`

	// Map contains query string parameters and their values.
	// It contains the same keys as QueryStringParameters, but if any parameter was repeated several times,
	// the list for it will contain all the passed values for this parameter.
	MultiValueQueryStringParameters map[string][]string `json:"multiValueQueryStringParameters"`

	// Meta information about request
	RequestContext RequestContext `json:"requestContext"`

	// Body of request in string format.
	// The data can be encoded in Base64 format (in this case IsBase64Encoded wil be set to true)
	Body            string `json:"body"`
	IsBase64Encoded bool   `json:"isBase64Encoded"`

	// FIXME: I don't see a description of these fields in the documentation.
	// During experiments, empty fields are returned
	URL string `json:"url"`

	// FIXME: I don't see a description of these fields in the documentation.
	// During experiments, empty fields are returned
	Params interface{} `json:"params"`

	// FIXME: I don't see a description of these fields in the documentation.
	// During experiments, empty fields are returned
	MultiValueParams interface{} `json:"multiValueParams"`

	// FIXME: I don't see a description of these fields in the documentation.
	// During experiments, empty fields are returned
	PathParams interface{} `json:"pathParams"`
}

func (req *Request) GetStringBody() (string, error) {
	var result string

	if req.IsBase64Encoded {
		bytes, err := base64.StdEncoding.DecodeString(req.Body)
		if err != nil {
			return "", err
		}
		result = string(bytes)
	} else {
		result = req.Body
	}
	return result, nil
}
