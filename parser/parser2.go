package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Gender string `json:"gender"`
	Name   Name `json:"name"`
}

type Name struct {
	Title string `json:"title"`
	FirstName string `json:"first"`
	LastName string `json:"last"`
}

func main()  {

	jsonResponse := `{"gender":"male","name":{"title":"mr","first":"wyatt","last":"garcia"}}`

	var user User

	err := json.Unmarshal([]byte(jsonResponse), &user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Gender: %s Title: %s First: %s Last: %s",
		user.Gender,
		user.Name.Title,
		user.Name.FirstName,
		user.Name.LastName)
}