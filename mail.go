package sg

import "encoding/json"

// Mail represents a SendGrid transactional mailer.
type Mail struct {
	From          string
	To            string
	TemplateId    string
	Substitutions map[string]string
}

// Implements the json.Marshaler interface.
func (m *Mail) MarshalJSON() ([]byte, error) {
	type h map[string]string
	type o map[string]interface{}

	// Don't send nil substitutions, the SendGrid API won't like it.
	substitutions := map[string]string{}
	if m.Substitutions != nil {
		substitutions = m.Substitutions
	}

	return json.Marshal(&struct {
		From             h      `json:"from"`
		Personalizations []o    `json:"personalizations"`
		Content          []h    `json:"content"`
		TemplateId       string `json:"template_id"`
	}{
		From:             h{"email": m.From},
		TemplateId:       m.TemplateId,
		Content:          []h{{"type": "text/html", "value": "<html><body></body></html>"}},
		Personalizations: []o{{"to": []h{{"email": m.To}}, "substitutions": substitutions}},
	})
}
