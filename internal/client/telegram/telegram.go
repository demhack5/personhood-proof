package telegram

type Client interface {
	SendMessage(userID int64, messageText string) error
	GetUpdates() (map[int64][]string, error)
}

type client struct {
}

var _ Client = &client{}

// mb use https://github.com/gotd/td ?
func NewClient() Client {
	return &client{}
}

func (c *client) SendMessage(userID int64, txt string) error {
	return nil
}

func (c *client) GetUpdates() (map[int64][]string, error) {
	return nil, nil
}
