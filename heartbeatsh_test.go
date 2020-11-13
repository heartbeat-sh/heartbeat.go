package heartbeat_go

import (
	"github.com/heartbeat-sh/heartbeat.go/heartbeatsh"
	"testing"
	"time"
)

const subdomain = "test"
const createTestHb = "go"
const nilTestHb = "go-nils"
const deleteTestHb = "go-delete"

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
	err = client.SendBeat(deleteTestHb, nil, nil)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
}

func TestDeleteBeat(t *testing.T) {
	c := heartbeatsh.NewClient(subdomain)
	err := c.DeleteBeat(deleteTestHb)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
}
