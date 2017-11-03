package sg

import (
	"bytes"
	"errors"
	"net/http"
)

// Client represents a SendGrid API v3 client.
type Client struct {
	APIKey  string
	APIURL  string
	Service Service
	Tracer  Tracer

	client http.Client
}

// Send sends a transactional mail as defined in the passed in Mail object.
func (c *Client) Send(mail *Mail) error {
	request, err := c.buildRequest(mail)
	if err != nil {
		return err
	}

	dumpRequest(c.Tracer, request)

	response, err := c.client.Do(request)
	if err != nil {
		return err
	}

	dumpResponse(c.Tracer, response)

	if response.StatusCode > 299 {
		return errors.New("bad request")
	}

	return nil
}

func (c *Client) buildRequest(mail *Mail) (*http.Request, error) {
	data, err := c.Service.Serialize(mail)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", c.APIURL, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", c.Service.Authorize(c.APIKey))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")

	return request, nil
}

// NewClient creates a new client with a SendGrid API key.
var NewClient = NewSendGridClient

// NewSendGridClient creates a new client with a SendGrid API key.
func NewSendGridClient(apiKey string) Sender {
	return &Client{
		APIKey:  apiKey,
		APIURL:  "https://api.sendgrid.com/v3/mail/send",
		Service: new(SendGridService),
	}
}

// NewSparkPostClient creates a new client with a SendGrid API key.
func NewSparkPostClient(apiKey string) Sender {
	return &Client{
		APIKey:  apiKey,
		APIURL:  "https://api.sparkpost.com/api/v1/transmissions?num_rcpt_errors=3",
		Service: new(SparkPostService),
	}
}
