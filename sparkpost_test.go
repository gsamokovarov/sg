package sg

import (
	"testing"

	"github.com/gsamokovarov/assert"
)

var sparkPost = &SparkPostService{}

func TestSparkPostService_WithoutSubscriptions(t *testing.T) {
	buf, err := sparkPost.Serialize(&Mail{
		TemplateID: "e0d26988-d1d7-41ad-b1eb-4c4b37125893",
		From:       "from@example.com",
		To:         "to@example.com",
	})
	assert.Nil(t, err)

	assert.Equal(t, `{"recipients":[{"address":"to@example.com"}],"content":{"template_id":"e0d26988-d1d7-41ad-b1eb-4c4b37125893"}}`, string(buf))
}

func TestSparkPostService_WithSubscriptions(t *testing.T) {
	buf, err := sparkPost.Serialize(&Mail{
		TemplateID:    "e0d26988-d1d7-41ad-b1eb-4c4b37125893",
		From:          "from@example.com",
		To:            "to@example.com",
		Substitutions: H{"SUB": "value"},
	})
	assert.Nil(t, err)

	assert.Equal(t, `{"recipients":[{"address":"to@example.com"}],"substitution_data":{"SUB":"value"},"content":{"template_id":"e0d26988-d1d7-41ad-b1eb-4c4b37125893"}}`, string(buf))
}
