package data

import (
	"net/url"
	"strings"
)

type TwilioRequest struct {
	SmsMessageSid string `json:"SmsMessageSid"`
	NumMedia      string `json:"NumMedia"`
	ProfileName   string `json:"ProfileName"`
	SmsSid        string `json:"SmsSid"`
	WaId          string `json:"WaId"`
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
		WaId:          m["WaId"],
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

func (tr TwilioRequest) ToUrlValues() url.Values {
	v := url.Values{}

	v.Set("SmsMessageSid", tr.SmsMessageSid)
	v.Set("NumMedia", tr.NumMedia)
	v.Set("ProfileName", tr.ProfileName)
	v.Set("SmsSid", tr.SmsSid)
	v.Set("WaId", tr.WaId)
	v.Set("SmsStatus", tr.SmsStatus)
	v.Set("Body", tr.Body)
	v.Set("To", tr.To)
	v.Set("NumSegments", tr.NumSegments)
	v.Set("MessageSid", tr.MessageSid)
	v.Set("AccountSid", tr.AccountSid)
	v.Set("From", tr.From)
	v.Set("ApiVersion", tr.ApiVersion)

	return v
}
