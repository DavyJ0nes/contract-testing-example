package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/davyj0nes/contract-testing-example/consumer"
)

// Client describes an http Client
type Client struct {
	BaseURL string
}

// New creates a client with url property
func New(url string) *Client {
	return &Client{
		BaseURL: url,
	}
}

// GetUser makes the http request to the URL to get a user
func (c *Client) GetUser(name string) (*consumer.User, error) {
	url := fmt.Sprintf("%s/users/%s", c.BaseURL, name)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	user := &consumer.User{}
	err = json.Unmarshal(body, user)

	return user, nil
}
