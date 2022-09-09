package client

import (
	"errors"
	"fmt"
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
	if res.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("client: invalid http response %d", res.StatusCode))
	}
	return nil
}
