package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler returns the given request body, but upper-case
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println(request)

	return events.APIGatewayProxyResponse{
		Body:       strings.ToUpper(request.Body),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
