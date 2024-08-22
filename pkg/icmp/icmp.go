package icmp

import (
	"fmt"
	"net"
	"time"
)

type Client struct {
	target  string
	timeout time.Duration
}

func NewClient(target string, timeout time.Duration) *Client {
	return &Client{target: target, timeout: timeout}
}

func (c *Client) SendEchoRequest() error {
	conn, err := net.Dial("ip4:icmp", c.target)
	if err != nil {
		return fmt.Errorf("failed to connect: %v", err)
	}
	defer conn.Close()

	// Construct ICMP Echo Request
	// Note: This is simplified. Real implementation should include ICMP header and checksum.
	msg := []byte{8, 0, 0, 0, 0, 0, 0, 0} // ICMP Echo Request (type 8, code 0)
	_, err = conn.Write(msg)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}

	return nil
}

func (c *Client) ReceiveEchoReply() ([]byte, error) {
	conn, err := net.Dial("ip4:icmp", c.target)
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %v", err)
	}
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(c.timeout))

	reply := make([]byte, 1500)
	n, err := conn.Read(reply)
	if err != nil {
		return nil, fmt.Errorf("failed to read reply: %v", err)
	}

	return reply[:n], nil
}
