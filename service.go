package sg

import "encoding/json"

// Service is the integration between the library and transactional mail
// service providers.
type Service interface {
	Authorize(string) string
	Serialize(*Mail) ([]byte, error)
}

// SendGridService serializes a mail for SendGrid API.
type SendGridService struct{}

// Authorize implements the Service interface.
func (*SendGridService) Authorize(key string) string { return "Bearer " + key }

// Serialize implements the Service interface.
func (*SendGridService) Serialize(m *Mail) ([]byte, error) {
	// Don't send nil substitutions, the SendGrid API won't like it and won't
	// give a good error message back.
	substitutions := H{}
	if m.Substitutions != nil {
		substitutions = m.Substitutions
	}

	return json.Marshal(&struct {
		From             H      `json:"from"`
		Personalizations []o    `json:"personalizations"`
		Content          []H    `json:"content"`
		TemplateID       string `json:"template_id"`
	}{
		From:             H{"email": m.From},
		TemplateID:       m.TemplateID,
		Content:          []H{{"type": "text/html", "value": "<html><body></body></html>"}},
		Personalizations: []o{{"to": []H{{"email": m.To}}, "substitutions": substitutions}},
	})
}

// SparkPostService serializes a mail for SparkPost API.
type SparkPostService struct{}

// Authorize implements the Service interface.
func (*SparkPostService) Authorize(key string) string { return key }

// Serialize implements the Service interface.
func (*SparkPostService) Serialize(m *Mail) ([]byte, error) {
	return json.Marshal(&struct {
		Recipients       []H `json:"recipients"`
		SubstitutionData H   `json:"substitution_data,omitempty"`
		Content          H   `json:"content"`
	}{
		Recipients:       []H{{"address": m.To}},
		Content:          H{"template_id": m.TemplateID},
		SubstitutionData: m.Substitutions,
	})
}
