package processing

import (
	"fmt"
	"github.com/m-to-n/channels-webhook-services/lambdas/whatsapp-twilio/data"
	"github.com/m-to-n/channels-webhook-services/utils"
)

func MessageHanler(twilioMessage *data.TwilioRequest) error {
	twilioMessageStr, err := utils.StructToString(twilioMessage)
	if err != nil {
		return err
	}
	fmt.Sprintf("MessageHanler called: %s", twilioMessageStr)
	return nil
}
