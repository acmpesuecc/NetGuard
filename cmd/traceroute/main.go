// cmd/traceroute/main.go
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/IshaanNene/NetGuard/pkg/traceroute" // Adjust the import path accordingly
)

func main() {
	target := flag.String("target", "", "The target host to trace")
	flag.Parse()

	if *target == "" {
		fmt.Fprintln(os.Stderr, "Error: No target host specified.")
		fmt.Fprintln(os.Stderr, "Usage: traceroute -target <host>")
		os.Exit(1)
	}

	tracerouteOutput, err := traceroute.Traceroute(*target)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running traceroute: %v\n", err)
		os.Exit(1)
	}

	formattedOutput := traceroute.ParseTracerouteOutput(tracerouteOutput)
	fmt.Println("Traceroute Output:")
	fmt.Println(formattedOutput)
}
