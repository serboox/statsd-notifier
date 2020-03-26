package stats

// ClientMocked client for mock
type ClientMocked struct{}

// NewClientMocked returns new mock structure instance
func NewClientMocked() *ClientMocked {
	return new(ClientMocked)
}

// Increment empty method for mock
func (c *ClientMocked) Increment(stat string) {}

// Close empty method for mock
func (c *ClientMocked) Close() {}
