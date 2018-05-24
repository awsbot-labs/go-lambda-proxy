# go-lambda-proxy

An example of how you could proxy http requests to a Lambda function.

Example use cases for this could be; testing your Lambda function locally before pushing to AWS, for integration testing and mocking API gateway.

**This is simply an example of how you might proxy HTTP traffic to a Lambda function. I wouldn't rely on this - instead, test your Lambdas with unit tests.**

```
  ________               _____________              ___________
 |        | > 1. http > |             | > 3. rpc > |           |
 | CLIENT |-------------| 2. PROXY 6. |------------| 4. LAMBDA |
 |________| < 7. http < |_____________| < 5. rpc < |___________|
```

1. The client makes an http request to the Lambda proxy.
2. The proxy creates an [`InvokeRequest`](https://github.com/aws/aws-lambda-go/blob/master/lambda/messages/messages.go#L16) using `http.Request` and [`APIGatewayProxyRequest`](https://github.com/aws/aws-lambda-go/blob/master/events/apigw.go#L6).
3. The proxy dials the Lambda and calls the `Function.Invoke` RPC with the `InvokeRequest`.
4. Lambda receives the `InvokeRequest` and the function runs.
5. Lambda responds to the RPC with an [`APIGatewayProxyResponse`](https://github.com/aws/aws-lambda-go/blob/master/events/apigw.go#L20).
6. The proxy parses the RPC response to an [`InvokeResponse`](https://github.com/aws/aws-lambda-go/blob/master/lambda/messages/messages.go#L27).
7. The proxy responds to the HTTP call; the response body is `InvokeResponse.Payload`.

## Getting started

1. Build the Lambda - `make build-lambda`
2. Build the proxy - `make build-proxy`
3. Start the Lambda - `make start-lambda`
4. Start the proxy - `make start-proxy`
5. Test your Lambda - `curl localhost:9898 -d "hello world" --silent | jq`

To use your own Lambda function, replace steps 1 and 3 with building and starting your own function, passing the `_LAMBDA_SERVER_PORT` environment variable. You can override the proxy ports with `PROXY_PORT` and `LAMBDA_PORT` flags.
