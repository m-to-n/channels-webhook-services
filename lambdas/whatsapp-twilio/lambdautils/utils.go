package lambdautils

import (
	"os"
	"strings"
)

func GetTwilioAuthKey(twilioMessageTo *string) string {
	return os.Getenv("TWILIO_AUTH_TOKEN_" + strings.Split(*twilioMessageTo, "+")[1]) // TODO: make more defensive :)
}

func GetSqsQueueUrl(twilioMessageTo *string) string {
	return os.Getenv("SQS_QUEUE_URL_" + strings.Split(*twilioMessageTo, "+")[1]) // TODO: make more defensive :)
}
