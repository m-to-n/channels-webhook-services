package processing

import (
	"fmt"
	"github.com/m-to-n/channels-webhook-services/lambdas/whatsapp-twilio/awssqs"
	"github.com/m-to-n/channels-webhook-services/lambdas/whatsapp-twilio/lambdautils"
	"github.com/m-to-n/channels-webhook-services/utils"
	whatsapp "github.com/m-to-n/common/channels/whatsapp-twilio"
)

func MessageHandler(twilioMessage *whatsapp.TwilioRequest) error {
	twilioMessageStr, err := utils.StructToString(twilioMessage)
	if err != nil {
		return err
	}
	fmt.Sprintf("MessageHanler called: %s", twilioMessageStr)

	queueUrl := lambdautils.GetSqsQueueUrl(&twilioMessage.To)

	sendResult, err := awssqs.SendChannelMessage(&queueUrl, twilioMessage)
	if err != nil {
		return err
	}

	sendResultStr, err := utils.StructToString(sendResult)
	if err != nil {
		return err
	}

	fmt.Sprintf("MessageHanler published sqs message: %s", sendResultStr)

	return nil
}
