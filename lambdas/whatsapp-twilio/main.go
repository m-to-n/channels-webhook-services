package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
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
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "OK!",
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

	if req.Body != "" {
		return handlerPost(ctx, req)
	} else {
		return handlerGet(ctx, req)
	}
}

func main() {
	lambda.Start(Handler)
}
