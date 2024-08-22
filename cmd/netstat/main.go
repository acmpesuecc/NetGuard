package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	showAll := flag.Bool("all", false, "Show all connections and listening ports")
	showTCP := flag.Bool("tcp", false, "Show TCP connections")
	showUDP := flag.Bool("udp", false, "Show UDP connections")
	flag.Parse()

	if *showAll {
		*showTCP = true
		*showUDP = true
	}

	if *showTCP {
		fmt.Println("TCP Connections:")
		listConnections("tcp")
	}

	if *showUDP {
		fmt.Println("UDP Connections:")
		listConnections("udp")
	}
}

func listConnections(networkType string) {
	var cmd *exec.Cmd
	var cmdArgs []string
	switch networkType {
	case "tcp":
		cmdArgs = []string{"-at"}
	case "udp":
		cmdArgs = []string{"-au"}
	default:
		log.Fatalf("Unsupported network type: %s", networkType)
	}
	cmd = exec.Command("netstat", cmdArgs...)

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error retrieving connections: %v", err)
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.Contains(line, networkType) || strings.Contains(line, networkType+"6") {
			fmt.Println(line)
		}
	}
}
