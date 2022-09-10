package main

import (
	client "erni93/form3-interview-accountapi/client"
	model "erni93/form3-interview-accountapi/models"
	"fmt"
)

func main() {
	client := client.WithDefaultConfig()
	fmt.Printf("Client available error %v\n", client.IsAvailable())

	accounts, err := client.GetAccounts()
	fmt.Printf("Get accounts: %v - error %v\n", accounts, err)

	country := "GB"
	newAccount := model.AccountData{
		Type:           "accounts",
		ID:             "ad27e262-9605-4b4b-a0e5-3003ea9cc4dc",
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Attributes: &model.AccountAttributes{
			Country: &country,
			BankID:  "400300",
			Bic:     "NWBKGB22",
			Name:    []string{"user1"},
		},
	}

	created, err := client.CreateAccount(newAccount)
	fmt.Printf("Create account: %v - error %v\n", created, err)
	if err != nil {
		panic(0)
	}

	account, err := client.GetAccount(created.Data.ID)
	fmt.Printf("Get account: %v - error %v\n", account, err)

	err = client.DeleteAccount(created.Data.ID, *created.Data.Version)
	fmt.Printf("Delete account error: %v\n", err)
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
