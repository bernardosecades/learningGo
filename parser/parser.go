package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Response struct {
	Status     string `json:"status"`
	Message    string `json:"message"`
}

func main()  {

	jsonResponse := `{"status":"success","message":"https:\/\/images.dog.ceo\/breeds\/bouvier\/n02106382_1623.jpg"}`

	var response Response

	err := json.Unmarshal([]byte(jsonResponse), &response)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Status: %s Message: %s", response.Status, response.Message)
}