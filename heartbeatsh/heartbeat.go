package heartbeatsh

import (
	"fmt"
	"net/http"
	"time"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{
		Timeout: time.Second * 5,
	}
}

type Client struct {
	proto     string
	host      string
	Subdomain string
}

// NewClient creates a new client for the server at {subdomain}.heartbeatsh.sh
func NewClient(subdomain string) Client {
	return Client{
		proto:     "https",
		host:      "heartbeat.sh",
		Subdomain: subdomain,
	}
}

// SendBeat sends a beat. To let the server choose timeouts, pass nil values to the timeout arguments.
// Returns any error that might be encountered with sending the http request.
func (c *Client) SendBeat(name string, warningTimeout *time.Duration, errorTimeout *time.Duration) error {
	var query string
	if warningTimeout != nil {
		query = fmt.Sprintf("?warning=%d", int(warningTimeout.Seconds()))
	}
	if errorTimeout != nil {
		sep := "&"
		if len(query) == 0 {
			sep = "?"
		}
		query = fmt.Sprintf("%s%serror=%d", query, sep, int(errorTimeout.Seconds()))
	}
	url := fmt.Sprintf("%v://%v.%v/beat/%v%v", c.proto, c.Subdomain, c.host, name, query)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	_, err = httpClient.Do(req)
	return err
}

// DeleteBeat deletes a beat.
// Returns any error that might be encountered with sending the http request.
func (c *Client) DeleteBeat(name string) error {
	url := fmt.Sprintf("%v://%v.%v/beat/%v", c.proto, c.Subdomain, c.host, name)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	_, err = httpClient.Do(req)
	return err
}
