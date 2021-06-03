package ycmodel

type Context struct {
	Awsrequestid       string  `json:"awsRequestId"`
	Requestid          string  `json:"requestId"`
	Invokedfunctionarn string  `json:"invokedFunctionArn"`
	Functionname       string  `json:"functionName"`
	Functionversion    string  `json:"functionVersion"`
	Memorylimitinmb    string  `json:"memoryLimitInMB"`
	Deadlinems         int64   `json:"deadlineMs"`
	Loggroupname       string  `json:"logGroupName"`
	Token              Token   `json:"token"`
	Request            Request `json:"_data"`
}

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}
