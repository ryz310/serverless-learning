package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       "Okay so your other function also executed successfully!",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
