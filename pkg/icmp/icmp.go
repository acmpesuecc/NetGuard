package icmp

import (
	"fmt"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
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

	msg := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   1,
			Seq:  1,
			Data: []byte("ping"),
		},
	}

	b, err := msg.Marshal(nil)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}

	_, err = conn.Write(b)
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
	_, err = conn.Read(reply)
	if err != nil {
		return nil, fmt.Errorf("failed to read reply: %v", err)
	}

	return reply, nil
}
