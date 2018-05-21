package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	tt := []struct {
		Name    string
		Input   string
		Output  string
		Success bool
	}{
		{
			"SuccessfulUpperCaseAlphaNumeric",
			"abcdefg12345",
			"ABCDEFG12345",
			true,
		},
	}

	for _, c := range tt {
		req := events.APIGatewayProxyRequest{Body: c.Input}
		res, err := Handler(req)

		if err != nil {
			t.Fatalf("%s failed: Handler returned an error: %v", c.Name, err)
		}

		if c.Success != (c.Output == res.Body) {
			t.Fatalf("%s failed: Expected %s but got %s", c.Name, c.Output, res.Body)
		}
	}
}
