build-lambda:
	cd lambda && GOOS=linux GOARCH=amd64 go build

test-lambda:
	cd lambda && go test --race

build-proxy:
	cd proxy && go build

test-proxy:
	cd proxy && go test --race
