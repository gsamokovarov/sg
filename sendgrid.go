package sg

import "encoding/json"

// SendGridService serializes a mail for SendGrid API.
type SendGridService struct{}

// Authorize implements the Service interface.
func (*SendGridService) Authorize(key string) string { return "Bearer " + key }

// Serialize implements the Service interface.
func (*SendGridService) Serialize(m *Mail) ([]byte, error) {
	// Don't send nil substitutions, the SendGrid API won't like it and there
	// won't be any decent error message back.
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

// NewSendGridClient creates a new client with a SendGrid API key.
func NewSendGridClient(apiKey string) Sender {
	return &Client{
		APIKey:  apiKey,
		APIURL:  "https://api.sendgrid.com/v3/mail/send",
		Service: new(SendGridService),
	}
}
