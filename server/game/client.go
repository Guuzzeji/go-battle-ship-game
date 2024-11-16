package client

type Client struct {
	ID   string
	Chat []string
}

func New(id string) *Client {
	return &Client{
		ID:   id,
		Chat: []string{},
	}
}

func (c *Client) AddMessage(msg string) {
	c.Chat = append(c.Chat, msg)
}

func (c *Client) GetLastMessage() string {
	if len(c.Chat) == 0 {
		return ""
	}

	return c.Chat[len(c.Chat)-1]
}
