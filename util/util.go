package util

import (
	"github.com/godruoyi/go-snowflake"
	"github.com/martinlindhe/base36"
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
