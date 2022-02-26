package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeStringToParams(t *testing.T) {
	t.Parallel()

	data := "SmsMessageSid=SMcd712b193a851d98e2e11ac6d3587c4c&NumMedia=0&ProfileName=Adam+Bezecny&SmsSid=SMcd712b193a851d98e2e11ac6d3587c4c&WaId=420123456789&SmsStatus=received&Body=Ahoy&To=whatsapp%3A%2B14155238886&NumSegments=1&MessageSid=SMcd712b193a851d98e2e11ac6d3587c4c&AccountSid=AC4f2a0dcdummy&From=whatsapp%3A%2B420608306441&ApiVersion=2010-04-01"

	arr, _ := DecodeStringToParams(data)
	assert.Equal(t, arr[0], "xxx") // TBD
}
