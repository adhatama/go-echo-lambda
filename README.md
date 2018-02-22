# Go Echo Lambda

An example of integration Go Echo with AWS Lambda.
We need to setup the AWS API Gateway to use Lambda Proxy Integration to execute this Lambda from API Gateway.

## Installation
- `go get github.com/aws/aws-lambda-go`
- `go get github.com/aws/aws-sdk-go`
- `go get github.com/labstack/echo`

## Build
To build for Lambda Deployment:
- `GOOS=linux GOARCH=amd64 go build`
- `zip main.zip echo-lambda`
- Set `LAMBDA` with `TRUE` in Environment variables in AWS Lambda console page.

To run it locally:
- `LAMBDA=TRUE go run main.go lambda_adapter.go`