package data

type TwilioRequest struct {
	SmsMessageSid string `json:"SmsMessageSid"`
	NumMedia      string `json:"NumMedia"`
	SmsSid        string `json:"SmsSid"`
	SmsStatus     string `json:"SmsStatus"`
	Body          string `json:"Body"`
	To            string `json:"To"`
	NumSegments   string `json:"NumSegments"`
	MessageSid    string `json:"MessageSid"`
	AccountSid    string `json:"AccountSid"`
	From          string `json:"From"`
	ApiVersion    string `json:"ApiVersion"`
}

type TwilioRequestValidationData struct {
	HeaderXTwilioSignature string
	HeaderXHost            string
	HeaderXPath            string
	AuthToken              string
}
