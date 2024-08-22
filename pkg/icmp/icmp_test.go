package icmp

import (
	"testing"
	"time"
)

func TestSendEchoRequest(t *testing.T) {
	client := NewClient("8.8.8.8", 2*time.Second)
	err := client.SendEchoRequest()
	if err != nil {
		t.Fatalf("SendEchoRequest failed: %v", err)
	}
}

func TestReceiveEchoReply(t *testing.T) {
	client := NewClient("8.8.8.8", 2*time.Second)
	_, err := client.ReceiveEchoReply()
	if err != nil {
		t.Fatalf("ReceiveEchoReply failed: %v", err)
	}
}
