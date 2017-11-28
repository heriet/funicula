package client

import (
	"github.com/heriet/funicula/nifcloud"
	"github.com/heriet/funicula/nifcloud/request"
)

// Client provides a nifcloud client
type Client struct {
	Config *nifcloud.Config
}

// New will return nifcloud client.
func New(cfg *nifcloud.Config) *Client {
	return &Client{
		Config: cfg,
	}
}

// NewRequest returns a new Request for the service API
func (c *Client) NewRequest(operation *request.Operation, params map[string]string, data interface{}) *request.Request {
	return request.New(c.Config, operation, params, data)
}
