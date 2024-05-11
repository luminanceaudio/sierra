package client

import "sierra/common/api"

func (c *Client) Shutdown() *api.Error {
	_, err := api.SendRequest[any, any](c.c, "", "POST", "app/shutdown", nil)
	if err != nil {
		return err
	}

	return nil
}
