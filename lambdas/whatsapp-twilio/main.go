package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/m-to-n/channels-webhook-services/lambdas/whatsapp-twilio/data"
	"github.com/m-to-n/channels-webhook-services/lambdas/whatsapp-twilio/lambdautils"
	"github.com/m-to-n/channels-webhook-services/lambdas/whatsapp-twilio/processing"
	"github.com/m-to-n/channels-webhook-services/lambdas/whatsapp-twilio/security"
	"github.com/m-to-n/channels-webhook-services/utils"
	"net/http"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func handlerGet(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "WA/Twilio webhook handler running here!",
	}, nil
}

func handlerPost(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	rawDecodedBytes, err := base64.StdEncoding.DecodeString(req.Body)

	if err != nil {
		return returnError("Error when decoding twilio payload", err, http.StatusBadRequest)
	}

	rawDecodedStr := string(rawDecodedBytes)

	fmt.Println("rawDecodedStr " + rawDecodedStr)

	paramsArr, err := utils.DecodeStringToParams(rawDecodedStr)

	if err != nil {
		return returnError("Error when parsing rawDecodedStr", err, http.StatusBadRequest)
	}

	twilioMessage := data.TwilioRequestFromArray(paramsArr)
	prettyString, err := utils.StructToPrettyString(twilioMessage)

	if err != nil {
		return returnError("Error prettyString", err, http.StatusBadRequest)
	}

	fmt.Println("prettyString " + *prettyString)

	if err != nil {
		return returnError("Error when parsing twilio payload", err, http.StatusBadRequest)
	}

	twilioAuthToken := lambdautils.GetTwilioAuthKey(&twilioMessage.To)

	err = security.ValidateIncomingRequest(
		"https://"+req.Headers["host"],
		twilioAuthToken,
		"/whatsapp-twilio",
		twilioMessage.ToUrlValues(),
		req.Headers["x-twilio-signature"],
	)

	if err != nil {
		return returnError("Invalid security header", err, http.StatusBadRequest)
	}

	fmt.Println("security check OK, processing now.")

	err = processing.MessageHandler(&twilioMessage)
	if err != nil {
		return returnError("Error when processing twilio payload", err, http.StatusInternalServerError)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "OK!",
	}, nil
}

func returnError(errorMsg string, err error, statusCode int) (events.APIGatewayProxyResponse, error) {
	fmt.Println(errorMsg + ": " + err.Error())
	fmt.Sprint(err)
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       errorMsg,
	}, nil
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	/**
	req.HTTPMethod not populated for some reason. we need to check for body instead! :(

	https://www.alexedwards.net/blog/serverless-api-with-go-and-aws-lambda
	https://github.com/aws/aws-lambda-go/issues/179
	https://www.bogotobogo.com/GoLang/GoLang_Serverless_WebAPI_with_AWS_Lambda.php
	*/

	fmt.Println("whatsapp-twilio lambda called " + req.Body)
	fmt.Println("HTTPMethod: " + req.HTTPMethod)
	fmt.Println("Headers: " + fmt.Sprint(req.Headers))

	if req.Body != "" {
		return handlerPost(ctx, req)
	} else {
		return handlerGet(ctx, req)
	}
}

func main() {
	lambda.Start(Handler)
}
