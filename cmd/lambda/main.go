package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"

	echoServer "jasonblanchard/depfootprint/pkg/echo"
)

var echoLambda *echoadapter.EchoLambdaV2

func init() {
	e, _ := echoServer.MakeEcho()
	echoLambda = echoadapter.NewV2(e)
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return echoLambda.Proxy(req)
}

func main() {
	lambda.Start(Handler)
}
