package sg

import (
	"os"
	"sync"
)

var (
	client   = NewSendGridClient(os.Getenv("SG_API_KEY"))
	clientMu = sync.Mutex{}
)

// Sender is the interface the clients follow.
type Sender interface {
	Send(*Mail) error
}

// Send sends a transactional mail as defined in the passed in Mail object.
func Send(mail *Mail) error {
	return client.Send(mail)
}

// Setup configures the default global SendGrid client.
func Setup(c Sender) {
	clientMu.Lock()
	defer clientMu.Unlock()

	client = c
}
