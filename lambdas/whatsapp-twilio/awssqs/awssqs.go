package awssqs

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	whatsapp "github.com/m-to-n/common/channels/whatsapp-twilio"
	"github.com/m-to-n/common/logging"
)

type SQSSendMessageAPI interface {
	SendMessage(ctx context.Context,
		params *sqs.SendMessageInput,
		optFns ...func(*sqs.Options)) (*sqs.SendMessageOutput, error)
}

// SendMsg sends a message to an Amazon SQS queue.
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If success, a SendMessageOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to SendMessage.
func SendMsg(c context.Context, api SQSSendMessageAPI, input *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	return api.SendMessage(c, input)
}

func SendChannelMessage(queueURL *string, twilioMessage *whatsapp.TwilioRequest) (*sqs.SendMessageOutput, error) {

	twilioMessageStr, err := logging.StructToString(twilioMessage)
	if err != nil {
		return nil, err
	}

	sMInput := &sqs.SendMessageInput{
		MessageAttributes: map[string]types.MessageAttributeValue{
			"Channel": {
				DataType:    aws.String("String"),
				StringValue: aws.String("Twilio/Whatsapp"),
			},
		},
		MessageGroupId:         &twilioMessage.From,
		MessageBody:            twilioMessageStr,
		MessageDeduplicationId: &twilioMessage.MessageSid,
		QueueUrl:               queueURL,
	}

	// TODO: move this into lambda initialization and execute only once!
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	client := sqs.NewFromConfig(cfg)

	return SendMsg(context.TODO(), client, sMInput)
}
