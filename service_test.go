package sg

import (
	"strings"
	"testing"
)

var sendGrid = &SendGridService{}

func TestSendGridService_WithoutSubscriptions(t *testing.T) {
	buf, err := sendGrid.Serialize(&Mail{
		TemplateID: "e0d26988-d1d7-41ad-b1eb-4c4b37125893",
		From:       "from@example.com",
		To:         "to@example.com",
	})
	if err != nil {
		t.Errorf("error: %v", err)
	}

	expected := `{"from":{"email":"from@example.com"},"personalizations":[{"substitutions":{},"to":[{"email":"to@example.com"}]}],"content":[{"type":"text/html","value":"\u003chtml\u003e\u003cbody\u003e\u003c/body\u003e\u003c/html\u003e"}],"template_id":"e0d26988-d1d7-41ad-b1eb-4c4b37125893"}`
	if got := string(buf); !strings.Contains(got, expected) {
		t.Errorf("Expected:\n%v, got:\n%v", expected, got)
	}
}

func TestSendGridService_WithSubscriptions(t *testing.T) {
	buf, err := sendGrid.Serialize(&Mail{
		TemplateID:    "e0d26988-d1d7-41ad-b1eb-4c4b37125893",
		From:          "from@example.com",
		To:            "to@example.com",
		Substitutions: H{"SUB": "value"},
	})
	if err != nil {
		t.Errorf("error: %v", err)
	}

	expected := `{"from":{"email":"from@example.com"},"personalizations":[{"substitutions":{"SUB":"value"},"to":[{"email":"to@example.com"}]}],"content":[{"type":"text/html","value":"\u003chtml\u003e\u003cbody\u003e\u003c/body\u003e\u003c/html\u003e"}],"template_id":"e0d26988-d1d7-41ad-b1eb-4c4b37125893"}`
	if got := string(buf); !strings.Contains(got, expected) {
		t.Errorf("Expected:\n%v, got:\n%v", expected, got)
	}
}

var sparkPost = &SparkPostService{}

func TestSparkPostService_WithoutSubscriptions(t *testing.T) {
	buf, err := sparkPost.Serialize(&Mail{
		TemplateID: "e0d26988-d1d7-41ad-b1eb-4c4b37125893",
		From:       "from@example.com",
		To:         "to@example.com",
	})
	if err != nil {
		t.Errorf("error: %v", err)
	}

	expected := `{"recipients":[{"address":"to@example.com"}],"content":{"template_id":"e0d26988-d1d7-41ad-b1eb-4c4b37125893"}}`
	if got := string(buf); !strings.Contains(got, expected) {
		t.Errorf("Expected:\n%v, got:\n%v", expected, got)
	}
}

func TestSparkPostService_WithSubscriptions(t *testing.T) {
	buf, err := sparkPost.Serialize(&Mail{
		TemplateID:    "e0d26988-d1d7-41ad-b1eb-4c4b37125893",
		From:          "from@example.com",
		To:            "to@example.com",
		Substitutions: H{"SUB": "value"},
	})
	if err != nil {
		t.Errorf("error: %v", err)
	}

	expected := `{"recipients":[{"address":"to@example.com"}],"substitution_data":{"SUB":"value"},"content":{"template_id":"e0d26988-d1d7-41ad-b1eb-4c4b37125893"}}`
	if got := string(buf); !strings.Contains(got, expected) {
		t.Errorf("Expected:\n%v, got:\n%v", expected, got)
	}
}
