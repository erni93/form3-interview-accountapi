package main

import (
	"encoding/json"
	client "erni93/form3-interview-accountapi/client"
	model "erni93/form3-interview-accountapi/models"
	"fmt"
	"os"
)

func main() {
	client := client.WithDefaultConfig()
	fmt.Printf("Client available %v/n", client.IsAvailable())

	client.GetAccounts()

	file, err := os.ReadFile("./samples/new-account-success-input.json")
	if err != nil {
		panic(err)
	}
	jsonText := string(file)
	input := model.NewAccountInput{}
	json.Unmarshal([]byte(jsonText), &input)
	data, err := client.CreateAccount(*input.Data)
	if err != nil {
		fmt.Printf("Error %s \n", err)
	}
	fmt.Printf("Final data %v \n", data)
}
