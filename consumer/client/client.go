package client

import (
	"io/ioutil"
	"net/http"
)

// Client describes an http Client
type Client struct {
	URL string
}

// New creates a client with url property
func New(url string) *Client {
	return &Client{
		URL: url,
	}
}

// Call makes the http request to the URL
func (c *Client) Call() (string, error) {
	res, err := http.Get(c.URL)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
