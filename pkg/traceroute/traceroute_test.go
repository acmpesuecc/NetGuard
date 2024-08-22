// pkg/traceroute/traceroute_test.go
package traceroute

import (
	"testing"
)

// MockTracerouteOutput simulates a sample output from the traceroute command.
const MockTracerouteOutput = `
traceroute to example.com (93.184.216.34), 30 hops max, 60 byte packets
 1  192.168.1.1 (192.168.1.1)  1.223 ms  1.272 ms  1.317 ms
 2  203.0.113.1 (203.0.113.1)  10.172 ms  10.231 ms  10.295 ms
 3  198.51.100.1 (198.51.100.1)  20.752 ms  20.804 ms  20.878 ms
 4  93.184.216.34 (93.184.216.34)  30.191 ms  30.252 ms  30.317 ms
`

func TestTraceroute(t *testing.T) {

	output, err := Traceroute("example.com")
	if err != nil {
		t.Fatalf("Traceroute failed: %v", err)
	}

	if len(output) == 0 {
		t.Fatal("Traceroute returned empty output")
	}
}

func TestParseTracerouteOutput(t *testing.T) {
	// Given mock output
	expectedOutput := MockTracerouteOutput
	actualOutput := ParseTracerouteOutput(MockTracerouteOutput)

	if actualOutput != expectedOutput {
		t.Errorf("ParseTracerouteOutput() = %v; want %v", actualOutput, expectedOutput)
	}
}
