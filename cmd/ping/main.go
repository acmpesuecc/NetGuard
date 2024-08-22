package main

import (
	"flag"
	"fmt"
	"github.com/IshaanNene/NetGuard/pkg/icmp"
	"log"
	"time"
)

func main() {
	target := flag.String("target", "8.8.8.8", "Target IP or domain")
	count := flag.Int("count", 4, "Number of ping requests")
	timeout := flag.Int("timeout", 1, "Timeout in seconds")
	flag.Parse()

	client := icmp.NewClient(*target, time.Duration(*timeout)*time.Second)

	for i := 0; i < *count; i++ {
		err := client.SendEchoRequest()
		if err != nil {
			log.Printf("Error sending request: %v", err)
			continue
		}

		reply, err := client.ReceiveEchoReply()
		if err != nil {
			log.Printf("Error receiving reply: %v", err)
			continue
		}

		fmt.Printf("Reply from %s: %s\n", *target, reply)
		time.Sleep(time.Second) // Delay between pings
	}
}
