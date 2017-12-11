package sg

import (
	"testing"

	"github.com/gsamokovarov/assert"
)

var sendGrid = &SendGridService{}

func TestSendGridService_WithoutSubscriptions(t *testing.T) {
	buf, err := sendGrid.Serialize(&Mail{
		TemplateID: "e0d26988-d1d7-41ad-b1eb-4c4b37125893",
		From:       "from@example.com",
		To:         "to@example.com",
	})
	assert.Nil(t, err)

	assert.Equal(t, `{"from":{"email":"from@example.com"},"personalizations":[{"substitutions":{},"to":[{"email":"to@example.com"}]}],"content":[{"type":"text/html","value":"\u003chtml\u003e\u003cbody\u003e\u003c/body\u003e\u003c/html\u003e"}],"template_id":"e0d26988-d1d7-41ad-b1eb-4c4b37125893"}`, string(buf))
}

func TestSendGridService_WithSubscriptions(t *testing.T) {
	buf, err := sendGrid.Serialize(&Mail{
		TemplateID:    "e0d26988-d1d7-41ad-b1eb-4c4b37125893",
		From:          "from@example.com",
		To:            "to@example.com",
		Substitutions: H{"SUB": "value"},
	})
	assert.Nil(t, err)

	assert.Equal(t, `{"from":{"email":"from@example.com"},"personalizations":[{"substitutions":{"SUB":"value"},"to":[{"email":"to@example.com"}]}],"content":[{"type":"text/html","value":"\u003chtml\u003e\u003cbody\u003e\u003c/body\u003e\u003c/html\u003e"}],"template_id":"e0d26988-d1d7-41ad-b1eb-4c4b37125893"}`, string(buf))
}
