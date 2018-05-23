package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/rpc"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda/messages"
)

var pp, lp int

func main() {
	flag.IntVar(&pp, "PROXY_PORT", 9898, "the proxy port")
	flag.IntVar(&lp, "LAMBDA_PORT", 8787, "the lambda port")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		req, err := createLambdaRequest(r)
		if err != nil {
			fmt.Fprintln(w, err)
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		res, err := parseLambdaResponse(req)
		if err != nil {
			fmt.Fprintln(w, err)
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		fmt.Fprintf(w, string(res.Payload))
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", pp), nil))
}

func createLambdaRequest(r *http.Request) (*messages.InvokeRequest, error) {
	b := r.Body
	defer b.Close()

	p, err := ioutil.ReadAll(b)
	if err != nil {
		return nil, fmt.Errorf("could not read the payload: %v", err)
	}

	pbs, err := json.Marshal(events.APIGatewayProxyRequest{
		HTTPMethod:      r.Method,
		Body:            string(p),
		IsBase64Encoded: false,
	})
	if err != nil {
		return nil, fmt.Errorf("could not marshall the payload: %v", err)
	}

	t := time.Now()
	return &messages.InvokeRequest{
		Payload:      pbs,
		RequestId:    "0",
		XAmznTraceId: "",
		Deadline: messages.InvokeRequest_Timestamp{
			Seconds: int64(t.Unix()),
			Nanos:   int64(t.Nanosecond()),
		},
		InvokedFunctionArn:    "",
		CognitoIdentityId:     "",
		CognitoIdentityPoolId: "",
		ClientContext:         []byte{},
	}, nil
}

func parseLambdaResponse(req *messages.InvokeRequest) (messages.InvokeResponse, error) {
	res := messages.InvokeResponse{}

	c, err := rpc.Dial("tcp", fmt.Sprintf(":%d", lp))
	if err != nil {
		return res, fmt.Errorf("could not create rpc client: %v", err)
	}
	defer c.Close()

	if err := c.Call("Function.Invoke", req, &res); err != nil {
		return res, fmt.Errorf("could not call lambda function: %v", err)
	}

	if res.Error != nil {
		return res, fmt.Errorf("lambda returned an error: %v", res.Error)
	}

	return res, nil
}
