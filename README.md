# channels-webhook-services

This repository holds set of AWS Lambda functions exposed via AWS API Gateway endpoints. Purpose of these functions is to:

* Receive webhook messages from respective messaging platform (e.g. WhatsApp provider (Twilio)) 
* Perform security validation of incoming request (e.g. validation of **x-twilio-signature** header and rejection of message not coming from Twilio) 
* Insert verified messages (i.e. verified **x-twilio-signature** + whitelisted **AccountSid** for Twilio WhatsApp) into AWS SQS queue for downstream processing 

Each channel (or channel and its provider should single channel be handled by multiple providers) will be implemented as separate lambda function

## Serverless architecture

We are using [Serverless Framework](https://www.serverless.com/) to make our life little bit easier. Quick getting started tutorial can be found [here](https://www.serverless.com/blog/framework-example-golang-lambda-support/).