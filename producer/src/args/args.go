package args

import "flag"

var NumberOfGroups *int
var BatchSize *int
var Interval *int
var OutputDir *string

// Init sets the given args for the program or fallsback on defaults
func Init() {
	NumberOfGroups = flag.Int("number-of-groups", 1, "number of groups to spawn")
	BatchSize = flag.Int("batch-size", 10, "batch size of the json output")
	Interval = flag.Int("interval", 10, "batch size of the json output")
	OutputDir = flag.String("output-directory", "./events", "the directory to output the files")
	flag.Parse()
}
