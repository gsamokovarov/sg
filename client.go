package sg

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

// Client represents a SendGrid API v3 client.
type Client struct {
	APIKey string
	APIURL string
	Tracer Tracer

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

func (c *Client) buildRequest(mail *Mail) (request *http.Request, err error) {
	buf := bytes.NewBuffer([]byte{})

	if err = json.NewEncoder(buf).Encode(mail); err != nil {
		return
	}

	if request, err = http.NewRequest("POST", c.APIURL, buf); err != nil {
		return
	}

	request.Header.Add("Authorization", "Bearer "+c.APIKey)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")

	return
}

// NewClient creates a new client with a SendGrid API key.
func NewClient(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
		APIURL: "https://api.sendgrid.com/v3/mail/send",
	}
}
