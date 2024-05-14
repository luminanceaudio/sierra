package client

import (
	"github.com/luminanceaudio/sierra/src/sierra/common/api"
	"net/http"
)

type Client struct {
	c *http.Client
}

func NewClient() *Client {
	return &Client{
		c: api.NewClient(),
	}
}
