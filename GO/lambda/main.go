package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

// Struct for the input the program expects from the client
type InfoEvent struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Age int `json:"age"`
}

// Struct for the output the server will send back to the client
type Response struct {
	Profile string `json:"profile"`
}

// Event handler, this function handles requests from clients
func HandleInfoEvent(event InfoEvent) (Response, error) {
	return Response{Profile: fmt.Sprintf("Their name is %s %s, they are %d ", event.Firstname, event.Lastname, event.Age)}, nil
}

func main() {
	lambda.Start(HandleInfoEvent)
}