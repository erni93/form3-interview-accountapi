package client

import (
	"bytes"
	"encoding/json"
	errorhandler "erni93/form3-interview-accountapi/errorhandler"
	model "erni93/form3-interview-accountapi/models"
	"fmt"
	"net/http"
	"strconv"
)

type ClientConfig struct {
	Hostname         string
	Version          string
	HealthEndpoint   string
	AccountsEndpoint string
	Client           *http.Client
}

type Client struct {
	Config ClientConfig
}

var (
	DefaultHostname         = "http://accountapi:8080"
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
			Client:           &http.Client{},
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
	res, err := c.Config.Client.Get(c.getURL(c.Config.HealthEndpoint))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = errorhandler.GetErrorResponse(res)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetAccounts() ([]*model.AccountData, error) {
	res, err := c.Config.Client.Get(c.getURL(c.Config.AccountsEndpoint))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = errorhandler.GetErrorResponse(res)
	if err != nil {
		return nil, err
	}

	var output model.DataAccountDataList
	err = json.NewDecoder(res.Body).Decode(&output)
	if err != nil {
		return nil, err
	}
	return output.Data, nil
}

func (c *Client) GetAccount(id string) (*model.AccountData, error) {
	res, err := c.Config.Client.Get(c.getURLWithId(c.Config.AccountsEndpoint, id))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = errorhandler.GetErrorResponse(res)
	if err != nil {
		return nil, err
	}

	var output model.DataAccountData
	err = json.NewDecoder(res.Body).Decode(&output)
	if err != nil {
		return nil, err
	}
	return output.Data, nil
}

func (c *Client) CreateAccount(data model.AccountData) (*model.AccountCreated, error) {
	input := model.DataAccountData{Data: &data}
	body, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	res, err := c.Config.Client.Post(c.getURL(c.Config.AccountsEndpoint), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	err = errorhandler.GetErrorResponse(res)
	if err != nil {
		return nil, err
	}

	var output model.DataAccountCreated
	err = json.NewDecoder(res.Body).Decode(&output)
	if err != nil {
		return nil, err
	}
	return output.Data, nil
}

func (c *Client) DeleteAccount(id string, version int64) error {
	req, err := http.NewRequest("DELETE", c.getURLWithId(c.Config.AccountsEndpoint, id), nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("version", strconv.FormatInt(version, 10))
	req.URL.RawQuery = q.Encode()

	res, err := c.Config.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	err = errorhandler.GetErrorResponse(res)
	if err != nil {
		return err
	}
	return nil
}
