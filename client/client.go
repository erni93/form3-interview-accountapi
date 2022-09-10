package client

import (
	"bytes"
	"encoding/json"
	errorhandler "erni93/form3-interview-accountapi/errorhandler"
	model "erni93/form3-interview-accountapi/models"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ClientConfig struct {
	Hostname         string
	Version          string
	HealthEndpoint   string
	AccountsEndpoint string
}

type Client struct {
	Config ClientConfig
}

var (
	DefaultHostname         = "http://localhost:8080"
	DefaultVersion          = "v1"
	DefaultHealthEndpoint   = "health"
	DefaultAccountsEndpoint = "organisation/accounts"
)

func WithDefaultConfig() Client {
	return Client{
		Config: ClientConfig{
			Hostname:         DefaultHostname,
			Version:          DefaultVersion,
			HealthEndpoint:   DefaultHealthEndpoint,
			AccountsEndpoint: DefaultAccountsEndpoint,
		},
	}
}

func (c *Client) getURL(endpoint string) string {
	return fmt.Sprintf("%s/%s/%s", c.Config.Hostname, c.Config.Version, endpoint)
}

func (c *Client) getURLWithId(endpoint string, id string) string {
	url := c.getURL(endpoint)
	return fmt.Sprintf("%s/%s/", url, id)
}

func (c *Client) IsAvailable() error {
	res, err := http.Get(c.getURL(c.Config.HealthEndpoint))
	if err != nil {
		return err
	}
	err = errorhandler.GetErrorResponse(res)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetAccounts() ([]model.AccountData, error) {
	res, err := http.Get(c.getURL(c.Config.AccountsEndpoint))
	if err != nil {
		return nil, err
	}
	err = errorhandler.GetErrorResponse(res)
	if err != nil {
		return nil, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return nil, err
	}
	fmt.Println(string(resBody))
	return make([]model.AccountData, 0), nil
}

func (c *Client) CreateAccount(data model.AccountData) (*model.AccountData, error) {
	input := model.NewAccountInput{Data: &data}
	body, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(c.getURL(c.Config.AccountsEndpoint), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	err = errorhandler.GetErrorResponse(res)
	if err != nil {
		return nil, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return nil, err
	}
	fmt.Println(string(resBody))
	return nil, nil
}
