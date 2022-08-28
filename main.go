package main

import (
	"flag"
	"fmt"
	"github.com/godruoyi/go-snowflake"
	"github.com/martinlindhe/base36"
	"os"
	"strings"
	"time"
)

type M map[string]interface{}

var RFC3339Milli = "2006-01-02T15:04:05.999Z"
var layout = "2006-01-02T15:04:05Z"

// Generate Flake IDs
func Generate(numberOfIds int, machineId uint16, startTime string, lowerCase bool) []string {
	if numberOfIds > 65535 {
		panic("Number of IDs cannot be greater than 65535")
	}

	st, _ := time.Parse(layout, startTime)
	snowflake.SetStartTime(st)
	snowflake.SetMachineID(machineId)

	ids := make([]string, numberOfIds)
	for i := 0; i < numberOfIds; i++ {
		var id = base36.Encode(snowflake.ID())
		if lowerCase {
			id = strings.ToLower(id)
		}
		ids[i] = id
	}

	return ids
}

// Parse Flake ID
func Parse(parse string, startTime string) []M {
	st, _ := time.Parse(layout, startTime)
	snowflake.SetStartTime(st)
	sid := snowflake.ParseID(base36.Decode(strings.ToUpper(parse)))

	// ex: MLBX46XNNK2
	var rows []M
	return append(rows,
		M{"ID": sid.ID},                // 82601872124276738
		M{"MACHINE_ID": sid.MachineID}, // 1023
		M{"SEQUENCE": sid.Sequence},    // 2
		M{"TIMESTAMP": st.UTC().UnixNano()/1e6 + int64(sid.Timestamp)}, // 1660689020982
		M{"DATE_TIME": sid.GenerateTime().Format(RFC3339Milli)},        // 2022-08-16T22:30:20.982Z
	)
}

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
		ids := Generate(*numberOfIds, uint16(*machineId), *startTime, *lowerCase)
		fmt.Print(strings.Join(ids, "\n"))
		os.Exit(0)
	}

	// Parse Flake ID
	parsed := Parse(*parse, *startTime)
	for _, row := range parsed {
		for key, val := range row {
			fmt.Printf("%s %v\n", key, val)
		}
	}
}
