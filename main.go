package main

import (
	"flag"
	"fmt"
	"github.com/serebro/flakeid/util"
	"os"
	"strings"
)

// Customized usage message
func usage() {
	fmt.Println("FlakeID - Generate Snowflake IDs")
	fmt.Println("Version: 1.0.0")
	fmt.Printf("Usage: %s [OPTIONS]\n", os.Args[0])
	flag.PrintDefaults()
}

// CLI
func main() {
	// Panic handler
	defer func() {
		if msg := recover(); msg != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Unexpected problem. %v\n", msg)
			os.Exit(1)
		}
	}()

	// Define command line options
	flag.Usage = usage
	var numberOfIds = flag.Int("n", 1, "Number of IDs to generate")
	var machineId = flag.Int("m", 0, "Machine ID. From 0 to 1023 (default 0)")
	var startTime = flag.String("s", "2022-01-01T00:00:00Z", "Epoch start time")
	var parse = flag.String("p", "", "Parse Flake ID. Ex: flakeid -p MLBX46XNNK2")
	var lowerCase = flag.Bool("l", false, "To lower-case. Ex: mlbx46xnnk2")

	// Parse command line options
	flag.Parse()

	// Generate Flake IDs
	if *parse == "" {
		ids := util.Generate(*numberOfIds, uint16(*machineId), *startTime, *lowerCase)
		fmt.Print(strings.Join(ids, "\n"))
		os.Exit(0)
	}

	// Parse Flake ID
	parsed := util.Parse(*parse, *startTime)
	for _, row := range parsed {
		for key, val := range row {
			fmt.Printf("%s %v\n", key, val)
		}
	}
}
