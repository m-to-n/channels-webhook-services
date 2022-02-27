package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeStringToParams(t *testing.T) {
	t.Parallel()

	data := "SmsMessageSid=SMcd712b193a851d98e2e11ac6d3587c4c&NumMedia=0&ProfileName=Adam+Bezecny&SmsSid=SMcd712b193a851d98e2e11ac6d3587c4c&WaId=420123456789&SmsStatus=received&Body=Foo%26Bar&To=whatsapp%3A%2B14155238886&NumSegments=1&MessageSid=SMcd712b193a851d98e2e11ac6d3587c4c&AccountSid=AC4f2a0dcdummy&From=whatsapp%3A%2B420608306441&ApiVersion=2010-04-01"

	real, _ := DecodeStringToParams(data)

	expected := [13]string{
		"SmsMessageSid=SMcd712b193a851d98e2e11ac6d3587c4c",
		"NumMedia=0",
		"ProfileName=Adam Bezecny",
		"SmsSid=SMcd712b193a851d98e2e11ac6d3587c4c",
		"WaId=420123456789",
		"SmsStatus=received",
		"Body=Foo&Bar",
		"To=whatsapp:+14155238886",
		"NumSegments=1",
		"MessageSid=SMcd712b193a851d98e2e11ac6d3587c4c",
		"AccountSid=AC4f2a0dcdummy",
		"From=whatsapp:+420608306441",
		"ApiVersion=2010-04-01",
	}

	// assert.Equal(t, reflect.DeepEqual(expected, real), true)
	for i := 0; i < 13; i++ {
		assert.Equal(t, real[i], expected[i])
	}
}
