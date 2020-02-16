package registry

import (
	"io"
	"net/http"
	"net/url"
)

// Client manages communication with the NPM registry.
type Client struct {
	BaseURL *url.URL
	client  http.Client
}

// NewRequest creates an NPM registry request.
func (c *Client) NewRequest(method, u string, body io.Reader) (*http.Request, error) {
	url, err := c.BaseURL.Parse(u)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return nil, err
	}

	// Request an abbreviated response with only the fields required to support installation
	req.Header.Set("Accept", "application/vnd.npm.install-v1+json; q=1.0, application/json; q=0.8, */*")

	return req, nil
}

// Do sends an HTTP request and returns an HTTP response, following policy (such as redirects, cookies, auth) as configured on the client.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
