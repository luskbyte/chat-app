package websocket

// Register registra o cliente no hub
func (c *Client) Register() {
	c.hub.register <- c
}

