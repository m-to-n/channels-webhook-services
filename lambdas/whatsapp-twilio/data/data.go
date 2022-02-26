package data

import "strings"

type TwilioRequest struct {
	SmsMessageSid string `json:"SmsMessageSid"`
	NumMedia      string `json:"NumMedia"`
	ProfileName   string `json:"ProfileName"`
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

func TwilioRequestFromArray(params []string) TwilioRequest {
	m := make(map[string]string)
	for _, param := range params {
		keyval := strings.Split(param, "=")
		key := keyval[0]
		val := keyval[1]
		m[key] = val
	}
	return TwilioRequest{
		SmsMessageSid: m["SmsMessageSid"],
		NumMedia:      m["NumMedia"],
		ProfileName:   m["ProfileName"],
		SmsSid:        m["SmsSid"],
		SmsStatus:     m["SmsStatus"],
		Body:          m["Body"],
		To:            m["To"],
		NumSegments:   m["NumSegments"],
		MessageSid:    m["MessageSid"],
		AccountSid:    m["AccountSid"],
		From:          m["From"],
		ApiVersion:    m["ApiVersion"],
	}
}

type TwilioRequestValidationData struct {
	HeaderXTwilioSignature string
	HeaderXHost            string
	HeaderXPath            string
	AuthToken              string
}
