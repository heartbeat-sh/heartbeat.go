package heartbeatsh

import (
	"testing"
	"time"
)

const subdomain = "test"
const createTestHb = "go"
const nilTestHb = "go-nils"
const typedNilTestHb = "go-typed-nils"
const deleteTestHb = "go-delete"
const goTimesTestHb = "go-times"

func TestNewClient(t *testing.T) {
	client := NewClient(subdomain)
	if client.Subdomain != subdomain {
		t.Errorf("Got %s, wanted %s", client.Subdomain, subdomain)
	}
}

func TestSendBeat(t *testing.T) {
	client := NewClient(subdomain)
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
	c := NewClient(subdomain)
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

func TestValidateTimeFormat(t *testing.T) {
	validTimes := []string{"09:00", "17:30", "00:00", "23:59"}
	invalidTimes := []string{"9:00", "17:5", "24:00", "12:60", "12:61", "12:5a"}

	for _, timeStr := range validTimes {
		if err := validateTimeFormat(timeStr); err != nil {
			t.Errorf("Expected valid time format for %s, got error %v", timeStr, err)
		}
	}

	for _, timeStr := range invalidTimes {
		if err := validateTimeFormat(timeStr); err == nil {
			t.Errorf("Expected error for invalid time format %s, got nil", timeStr)
		}
	}
}

func TestDeleteBeat(t *testing.T) {
	c := NewClient(subdomain)
	err := c.DeleteBeat(deleteTestHb)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
}
