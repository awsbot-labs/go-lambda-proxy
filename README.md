# go-lambda-proxy

An example of how you could proxy HTTP requests to a Lambda function locally.

**This is simply an example of how you might proxy HTTP traffic to RPC in the context of Lambda. I wouldn't recommend doing this - instead, test your Lambdas with unit tests.**

## Getting started

1. Build the Lambda - `make build-lambda`
2. Build the proxy - `make build-proxy`
3. Start the Lambda - `make start-lambda`
4. Start the proxy - `make start-proxy`
5. Test your Lambda - `curl localhost:9898 -d '{"Body": "something"}' --silent | jq`

To use your own Lambda function, replace steps 1 and 3 with building and starting your own function, passing the `_LAMBDA_SERVER_PORT` environment variable. You can override the proxy ports with `PROXY_PORT` and `LAMBDA_PORT` flags.
