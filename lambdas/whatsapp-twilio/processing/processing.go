package processing

import (
	"fmt"
	"github.com/m-to-n/channels-webhook-services/lambdas/whatsapp-twilio/awssqs"
	"github.com/m-to-n/channels-webhook-services/lambdas/whatsapp-twilio/lambdautils"
	whatsapp "github.com/m-to-n/common/channels/whatsapp-twilio"
	"github.com/m-to-n/common/logging"
)

func MessageHandler(twilioMessage *whatsapp.TwilioRequest) error {
	twilioMessageStr, err := logging.StructToString(twilioMessage)
	if err != nil {
		return err
	}
	fmt.Sprintf("MessageHanler called: %s", twilioMessageStr)

	queueUrl := lambdautils.GetSqsQueueUrl(&twilioMessage.To)

	sendResult, err := awssqs.SendChannelMessage(&queueUrl, twilioMessage)
	if err != nil {
		return err
	}

	sendResultStr, err := logging.StructToString(sendResult)
	if err != nil {
		return err
	}

	fmt.Sprintf("MessageHanler published sqs message: %s", sendResultStr)

	return nil
}
