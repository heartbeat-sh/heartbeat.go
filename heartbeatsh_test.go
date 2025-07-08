package heartbeat_go

import (
	"testing"
	"time"

	"github.com/heartbeat-sh/heartbeat.go/heartbeatsh"
)

const subdomain = "test"
const createTestHb = "go"
const nilTestHb = "go-nils"
const typedNilTestHb = "go-typed-nils"
const deleteTestHb = "go-delete"
const goTimesTestHb = "go-times"

func TestNewClient(t *testing.T) {
	client := heartbeatsh.NewClient(subdomain)
	if client.Subdomain != subdomain {
		t.Errorf("Got %s, wanted %s", client.Subdomain, subdomain)
	}
}

func TestSendBeat(t *testing.T) {
	client := heartbeatsh.NewClient(subdomain)
	minute := time.Minute
	hour := time.Hour
	err := client.SendBeat(createTestHb, &minute, &hour)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	err = client.SendBeat(nilTestHb, nil, nil)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	var noTimeout *time.Duration
	err = client.SendBeat(typedNilTestHb, noTimeout, noTimeout)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	err = client.SendBeat(deleteTestHb, nil, nil)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
}

func TestSendBeatWithTimes(t *testing.T) {
	c := heartbeatsh.NewClient(subdomain)
	minute := time.Minute
	hour := time.Hour

	err := c.SendBeatWithTimes(createTestHb, &minute, &hour, "09:00", "17:00")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	err = c.SendBeatWithTimes(nilTestHb, nil, nil, "", "")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	var noTimeout *time.Duration
	err = c.SendBeatWithTimes(typedNilTestHb, noTimeout, noTimeout, "", "")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	err = c.SendBeatWithTimes(deleteTestHb, nil, nil, "", "")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	err = c.SendBeatWithTimes(goTimesTestHb, &minute, &hour, "09:00", "17:00")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	err = c.SendBeatWithTimes(goTimesTestHb, &minute, &hour, "09", "17")
	if err == nil {
		t.Errorf("Expected error for invalid time format, got nil")
	}
}

func TestDeleteBeat(t *testing.T) {
	c := heartbeatsh.NewClient(subdomain)
	err := c.DeleteBeat(deleteTestHb)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
}
