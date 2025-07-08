package heartbeatsh

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
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
	return c.SendBeatWithTimes(name, warningTimeout, errorTimeout, "", "")
}

// SendBeatWithTimes sends a beat. To let the server choose timeouts, pass nil values to the timeout arguments with optional start and end times.
// startTime and endTime should be in "HH:MM" format (e.g., "09:00", "17:00"). Pass empty strings to omit.
// Providing a startTime or endTime define a time window for when the beat should be actively checked.
// Returns any error that might be encountered with sending the http request.
func (c *Client) SendBeatWithTimes(name string, warningTimeout *time.Duration, errorTimeout *time.Duration, startTime string, endTime string) error {

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
	if startTime != "" {
		if err := validateTimeFormat(startTime); err != nil {
			return fmt.Errorf("invalid startTime: %w", err)
		}

		sep := "&"
		if len(query) == 0 {
			sep = "?"
		}
		query = fmt.Sprintf("%s%sstartTime=%s", query, sep, startTime)
	}
	if endTime != "" {
		if err := validateTimeFormat(endTime); err != nil {
			return fmt.Errorf("invalid endTime: %w", err)
		}

		sep := "&"
		if len(query) == 0 {
			sep = "?"
		}
		query = fmt.Sprintf("%s%sendTime=%s", query, sep, endTime)
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

// validateTimeFormat validates that a time string is in "HH:MM" format
// Returns an error if the format is invalid or the time values are out of range
func validateTimeFormat(timeStr string) error {
	timeRegex := regexp.MustCompile(`^([0-1][0-9]|2[0-3]):([0-5][0-9])$`)
	if !timeRegex.MatchString(timeStr) {
		return errors.New("time must be in HH:MM format (e.g., '09:00', '17:30')")
	}

	parts := regexp.MustCompile(`:`).Split(timeStr, 2)
	if len(parts) != 2 {
		return errors.New("time must be in HH:MM format")
	}

	hours := parts[0]
	minutes := parts[1]

	h, err := strconv.Atoi(hours)
	if err != nil || h < 0 || h > 23 {
		return errors.New("hour must be between 00 and 23")
	}

	m, err := strconv.Atoi(minutes)
	if err != nil || m < 0 || m > 59 {
		return errors.New("minute must be between 00 and 59")
	}

	return nil
}
