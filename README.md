## The problem described below is fixed! I was responding to the ping with a capital 'T' in "type". The code has been corrected.

I want to set up a new Discord bot to interact via an HTTP endpoint, but I can't get it to work.
(aws lambda in golang, exposed through the aws api gateway)

Here's the main function :
```go
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
```

In summary, it tests the validity of the header, then checks if it's a PING.

When I set up the Discord bot to use this api (in "INTERACTIONS ENDPOINT URL"), it tells me that the endpoint couldn't be verified. The log indicates an authorization error on the second request.

Yet the log shows me responding correctly to the PING on the first request (so the header validation passed).

However, I use exactly the same code on another project that works perfectly, so it can't be much!

The full trace :
```
START RequestId: 590104c1-4353-4153-8708-614c6afcdbbc Version: $LATEST
start
test ping
is ping !
ping
END RequestId: 590104c1-4353-4153-8708-614c6afcdbbc
REPORT RequestId: 590104c1-4353-4153-8708-614c6afcdbbc	Duration: 36.13 ms	Billed Duration: 37 ms	Memory Size: 128 MB	Max Memory Used: 29 MB	Init Duration: 97.97 ms	
START RequestId: b9e94769-0b0d-421b-91c3-72748f333567 Version: $LATEST
start
StatusUnauthorized
END RequestId: b9e94769-0b0d-421b-91c3-72748f333567
REPORT RequestId: b9e94769-0b0d-421b-91c3-72748f333567	Duration: 18.34 ms	Billed Duration: 19 ms	Memory Size: 128 MB	Max Memory Used: 30 MB	
```

The full code is here : https://github.com/DavidVergison/discord-bot-error

Do you have any idea what I'm forgetting ?