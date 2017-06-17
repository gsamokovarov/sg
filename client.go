package sg

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httputil"
)

// Client represents a SendGrid API v3 client.
type Client struct {
	ApiKey string
	ApiUrl string
	Tracer Tracer

	client http.Client
}

// Send sends a transactional mail as defined in the passed in Mail object.
func (c *Client) Send(mail *Mail) error {
	request, err := c.buildRequest(mail)
	if err != nil {
		return err
	}

	c.dumpRequest(request)

	response, err := c.client.Do(request)
	if err != nil {
		return err
	}

	c.dumpResponse(response)

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

	if request, err = http.NewRequest("POST", c.ApiUrl, buf); err != nil {
		return
	}

	request.Header.Add("Authorization", "Bearer "+c.ApiKey)
	request.Header.Add("Content-Type", "application/json")

	return
}

func (c *Client) dumpRequest(request *http.Request) {
	if c.Tracer == nil {
		return
	}

	if dump, err := httputil.DumpRequest(request, true); err == nil {
		c.Tracer.Printf("\n%s\n", dump)
	}
}

func (c *Client) dumpResponse(response *http.Response) {
	if c.Tracer == nil {
		return
	}

	if dump, err := httputil.DumpResponse(response, true); err == nil {
		c.Tracer.Printf("\n%s\n", dump)
	}
}

// NewClient creates a new client with a SendGrid API key.
func NewClient(apiKey string) *Client {
	return &Client{
		ApiKey: apiKey,
		ApiUrl: "https://api.sendgrid.com/v3/mail/send",
	}
}
