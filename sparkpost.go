package sg

import "encoding/json"

// SparkPostService serializes a mail for SparkPost API.
type SparkPostService struct{}

// Authorize implements the Service interface.
func (*SparkPostService) Authorize(key string) string { return key }

// Serialize implements the Service interface.
func (*SparkPostService) Serialize(m *Mail) ([]byte, error) {
	content := o{"template_id": m.TemplateID}
	if m.Attachments != nil {
		content["attachments"] = m.Attachments
	}

	return json.Marshal(&struct {
		Recipients       []H `json:"recipients"`
		SubstitutionData H   `json:"substitution_data,omitempty"`
		Content          o   `json:"content"`
	}{
		Recipients:       []H{{"address": m.To}},
		Content:          content,
		SubstitutionData: m.Substitutions,
	})
}

// NewSparkPostClient creates a new client with a SparkPost API key.
func NewSparkPostClient(apiKey string, tracers ...Tracer) Sender {
	return &Client{
		APIKey:  apiKey,
		APIURL:  "https://api.sparkpost.com/api/v1/transmissions?num_rcpt_errors=3",
		Service: new(SparkPostService),
		Tracer:  composedTracer{tracers},
	}
}
