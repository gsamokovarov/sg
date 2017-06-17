package sg

import "encoding/json"

type o map[string]interface{}

// H is a shortcut for map[string]string. In Go 1.9 this will become a type
// alias.
type H map[string]string

// Mail represents a SendGrid transactional mailer.
type Mail struct {
	From          string
	To            string
	TemplateID    string
	Substitutions H
}

// MarshalJSON implements the json.Marshaler interface.
func (m *Mail) MarshalJSON() ([]byte, error) {
	// Don't send nil substitutions, the SendGrid API won't like it.
	substitutions := map[string]string{}
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
		TemplateID:       m.TemplateId,
		Content:          []H{{"type": "text/html", "value": "<html><body></body></html>"}},
		Personalizations: []o{{"to": []H{{"email": m.To}}, "substitutions": substitutions}},
	})
}
