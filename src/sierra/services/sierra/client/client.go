package client

import (
	"net/http"
	"sierra/common/api"
)

type Client struct {
	c *http.Client
}

func NewClient() *Client {
	return &Client{
		c: api.NewClient(),
	}
}
