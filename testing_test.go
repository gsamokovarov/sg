package sg

import (
	"testing"

	"github.com/gsamokovarov/assert"
)

func TestTestingClient(t *testing.T) {
	c := NewTestingClient()

	_, err := c.Last()
	assert.Error(t, err, "sg: inbox is empty")

	mail := &Mail{
		From: "me@example.com",
		To:   "you@example.com",
	}

	err = c.Send(mail)
	assert.Nil(t, err)

	m, err := c.Last()
	assert.Nil(t, err)

	assert.Equal(t, m, mail)
}
