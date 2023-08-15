package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	rawBody := request.Body
	fmt.Println("start")

	// validate nacl header
	err := validator.verifyRequest(
		request.Headers["x-signature-ed25519"],
		request.Headers["x-signature-timestamp"],
		request.Body,
	)
	if err != nil {
		fmt.Println("StatusUnauthorized")
		return formatUnauthorizedResponse()
	}

	discordRequest := DiscordRequestDto{}
	err = json.Unmarshal([]byte(rawBody), &discordRequest)

	// validate ping
	fmt.Println("test ping")
	if validator.IsPing(discordRequest.Type) {
		fmt.Println("is ping !")
		return formatPingResponse()
	}

	return formatResponse("ok")
}

func formatPingResponse() (events.APIGatewayProxyResponse, error) {
	fmt.Println("ping")
	json, _ := json.Marshal(
		struct {
			Type int
		}{
			Type: 1,
		},
	)

	return events.APIGatewayProxyResponse{
		Body:       string(json),
		StatusCode: 200,
	}, nil
}

func formatResponse(msg string) (events.APIGatewayProxyResponse, error) {
	rawResponse := DiscordResponseDto{
		Type: 4,
		Data: DiscordDataResponseDto{
			Content: msg,
		},
	}

	jsonResponse, _ := json.Marshal(rawResponse)

	return events.APIGatewayProxyResponse{
		Body:       string(jsonResponse),
		StatusCode: 200,
	}, nil
}

func formatUnauthorizedResponse() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusUnauthorized,
	}, nil
}

var validator DiscordCommandValidator

func main() {
	validator = DiscordCommandValidator{
		publicKey: "178d3620f3d84a470ecba469d25f8aeb7f4e674065c3482e34df257e9a62f832",
	}
	lambda.Start(handler)
}
