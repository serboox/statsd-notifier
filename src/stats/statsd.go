package stats

import statsd "github.com/etsy/statsd/examples/go"

// ClientInterface interface for methods used in statsd
type ClientInterface interface {
	Increment(stat string)
	Close()
}

// Client statsd client structure
type Client struct {
	client *statsd.StatsdClient
}

// NewClient creates new statsd client
func NewClient(client *statsd.StatsdClient) *Client {
	return &Client{
		client: client,
	}
}

// Increment increments one stat counter without sampling
func (c *Client) Increment(stat string) {
	c.client.Increment(stat)
}

// Close method to close udp connection
func (c *Client) Close() {
	c.client.Close()
}
