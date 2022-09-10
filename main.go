package main

import (
	client "erni93/form3-interview-accountapi/client"
	"fmt"
)

func main() {
	client := client.WithDefaultConfig()
	fmt.Printf("Client available %v/n", client.IsAvailable())

	client.GetAccounts()

	err := client.DeleteAccount("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", int64(0))
	fmt.Printf("Result %v \n", err)
	/*

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
	*/
}
