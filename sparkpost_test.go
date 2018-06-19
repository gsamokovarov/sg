package sg

import (
	"testing"

	"github.com/gsamokovarov/assert"
)

var sparkPost = &SparkPostService{}

func TestSparkPostService(t *testing.T) {
	t.Run("without subscriptions", func(t *testing.T) {
		buf, err := sparkPost.Serialize(&Mail{
			TemplateID: "e0d26988-d1d7-41ad-b1eb-4c4b37125893",
			To:         "to@example.com",
		})
		assert.Nil(t, err)

		assert.Equal(t, `{"recipients":[{"address":"to@example.com"}],"content":{"template_id":"e0d26988-d1d7-41ad-b1eb-4c4b37125893"}}`, string(buf))
	})

	t.Run("with subscriptions", func(t *testing.T) {
		buf, err := sparkPost.Serialize(&Mail{
			TemplateID:    "e0d26988-d1d7-41ad-b1eb-4c4b37125893",
			To:            "to@example.com",
			Substitutions: H{"SUB": "value"},
		})
		assert.Nil(t, err)

		assert.Equal(t, `{"recipients":[{"address":"to@example.com"}],"substitution_data":{"SUB":"value"},"content":{"template_id":"e0d26988-d1d7-41ad-b1eb-4c4b37125893"}}`, string(buf))
	})

	t.Run("with inline template", func(t *testing.T) {
		buf, err := sparkPost.Serialize(&Mail{
			From:           "genadi@example.com",
			FromName:       "Genadi Samokovarov",
			To:             "to@example.com",
			Subject:        "Test Email",
			TemplateInline: `Testing 1, 2, 3...`,
			Substitutions:  H{"SUB": "value"},
		})
		assert.Nil(t, err)

		assert.Equal(t, `{"recipients":[{"address":"to@example.com"}],"substitution_data":{"SUB":"value"},"content":{"from":{"email":"genadi@example.com","name":"Genadi Samokovarov"},"html":"Testing 1, 2, 3...","subject":"Test Email"}}`, string(buf))
	})
}
