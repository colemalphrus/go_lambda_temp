package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)


type Data struct {
	ID string
}

type Response struct {
	ID string `json:"id"`
	Msg string `json:"msg"`
	RandomKey int `json:"key"`
}
 

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Recived Request:", request.Body)

	var d Data
	json.Unmarshal([]byte(request.Body), &d)

	resp := Response{
		ID: d.ID,
		Msg: fmt.Sprintf(" the users id is %s", d.ID),
		RandomKey: 50000,
	}
	r, _ := json.Marshal(resp)

	res := events.APIGatewayProxyResponse {
		Body: string(r),
		StatusCode: 200,
	}
	return res, nil
}

func main()  {
	lambda.Start(Handler)
}