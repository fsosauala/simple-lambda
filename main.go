package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, data events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("hello", data.Body)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: `{"hello":"my baby"}`,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
