package sg

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestMail_MarshalJSON_WithoutSubscriptions(t *testing.T) {
	mail := &Mail{
		TemplateID: "e0d26988-d1d7-41ad-b1eb-4c4b37125893",
		From:       "from@example.com",
		To:         "to@example.com",
	}

	buf := &bytes.Buffer{}

	if err := json.NewEncoder(buf).Encode(mail); err != nil {
		t.Errorf("error: %v", err)
	}

	expected := `{"from":{"email":"from@example.com"},"personalizations":[{"substitutions":{},"to":[{"email":"to@example.com"}]}],"content":[{"type":"text/html","value":"\u003chtml\u003e\u003cbody\u003e\u003c/body\u003e\u003c/html\u003e"}],"template_id":"e0d26988-d1d7-41ad-b1eb-4c4b37125893"}`
	if got := buf.String(); !strings.Contains(got, expected) {
		t.Errorf("Expected:\n%v, got:\n%v", expected, got)
	}
}

func TestMail_MarshalJSON_WithSubscriptions(t *testing.T) {
	mail := &Mail{
		TemplateID:    "e0d26988-d1d7-41ad-b1eb-4c4b37125893",
		From:          "from@example.com",
		To:            "to@example.com",
		Substitutions: H{"SUB": "value"},
	}

	buf := &bytes.Buffer{}

	if err := json.NewEncoder(buf).Encode(mail); err != nil {
		t.Errorf("error: %v", err)
	}

	expected := `{"from":{"email":"from@example.com"},"personalizations":[{"substitutions":{"SUB":"value"},"to":[{"email":"to@example.com"}]}],"content":[{"type":"text/html","value":"\u003chtml\u003e\u003cbody\u003e\u003c/body\u003e\u003c/html\u003e"}],"template_id":"e0d26988-d1d7-41ad-b1eb-4c4b37125893"}`
	if got := buf.String(); !strings.Contains(got, expected) {
		t.Errorf("Expected:\n%v, got:\n%v", expected, got)
	}
}
