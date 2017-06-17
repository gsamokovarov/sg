package sg

import (
	"os"
	"sync"
)

var (
	client = &Client{
		APIKey: os.Getenv("SG_API_KEY"),
		APIURL: "https://api.sendgrid.com/v3/mail/send",
	}

	clientMu = sync.Mutex{}
)

// Send sends a transactional mail as defined in the passed in Mail object.
func Send(mail *Mail) error {
	return client.Send(mail)
}

// Setup configures the default global SendGrid client.
func Setup(c *Client) {
	clientMu.Lock()
	defer clientMu.Unlock()

	client = c
}
