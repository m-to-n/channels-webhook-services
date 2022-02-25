package security

import "github.com/m-to-n/channels-webhook-services/lambdas/whatsapp-twilio/data"

func VerifySiganture(data data.TwilioRequestValidationData) bool {
	return true
}
