// pkg/traceroute/traceroute.go
package traceroute

import (
	"os/exec"
)

// Traceroute runs the traceroute command and returns the output.
func Traceroute(target string) (string, error) {
	// Run the traceroute command
	cmd := exec.Command("traceroute", target)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// ParseTracerouteOutput parses the traceroute output and returns a formatted string.
func ParseTracerouteOutput(output string) string {
	// For simplicity, we're just returning the output as is.
	// You can add more complex parsing logic here if needed.
	return output
}
