package log_server

type client struct {
}

func NewClient() (c *client, err error) {
	c = &client{}
	return c, nil
}
