package main

import (
	Client "erni93/form3-interview-accountapi/client"
	"fmt"
)

func main() {
	client := Client.WithDefaultConfig()
	fmt.Printf("Client available %v", client.IsAvailable())
}
