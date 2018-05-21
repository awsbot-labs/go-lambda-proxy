test-lambda:
	cd lambda && go test --race

build-lambda:
	cd lambda && GOOS=linux GOARCH=amd64 go build

start-lambda:
	_LAMBDA_SERVER_PORT=8787 ./lambda/lambda

build-proxy:
	cd proxy && go build

start-proxy:
	./proxy/proxy