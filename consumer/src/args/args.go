package args

import "flag"

var InputDir *string

// Init sets the given args for the program or fallsback on defaults
func Init() {
	InputDir = flag.String("input-directory", "./events", "the directory of the input files")
	flag.Parse()
}
