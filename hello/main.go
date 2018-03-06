package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       "Go Serverless v1.0! Your function executed successfully!",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
