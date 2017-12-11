package sg

import (
	"errors"
	"sync"
)

// TestingClient is a client that aggregates the sent mails in it's inbox. Use
// it to introspect mail during testing.
type TestingClient struct {
	Inbox []*Mail
	mu    sync.Mutex
}

// Send sends a transactional mail to the testing client inbox field.
func (c *TestingClient) Send(mail *Mail) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Inbox = append(c.Inbox, mail)

	return nil
}

// Last returns the last sent email.
func (c *TestingClient) Last() (*Mail, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.Inbox) == 0 {
		return nil, errors.New("sg: inbox is empty")
	}

	return c.Inbox[len(c.Inbox)-1], nil
}

// Empty empties the client inbox.
func (c *TestingClient) Empty() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c != nil {
		c.Inbox = make([]*Mail, 0)
	}
}

// NewTestingClient creates a new testing client.
func NewTestingClient() *TestingClient {
	return new(TestingClient)
}
