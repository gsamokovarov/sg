package sg

import (
	"os"
	"sync"
)

var (
	client   = NewSendGridClient(os.Getenv("SG_API_KEY"))
	clientMu = sync.Mutex{}
)

// Service is the integration between the library and transactional mail
// service providers.
type Service interface {
	Authorize(string) string
	Serialize(*Mail) ([]byte, error)
}

// Sender is the interface the clients follow. The sender can send a mail.
type Sender interface {
	Send(*Mail) error
}

// Send sends a transactional mail as defined in the passed in Mail object.
func Send(mail *Mail) error {
	return client.Send(mail)
}

// Setup configures the default global sg client.
func Setup(c Sender) {
	clientMu.Lock()
	defer clientMu.Unlock()

	client = c
}
