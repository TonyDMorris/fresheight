package main

import (
	"os"

	"github.com/tonydmorris/fresh8-producer/src/args"
	"github.com/tonydmorris/fresh8-producer/src/batch"
	"github.com/tonydmorris/fresh8-producer/src/contrive"
	"github.com/tonydmorris/fresh8-producer/src/logger"
)

func main() {
	// init program args
	args.Init()
	err := os.MkdirAll(*args.OutputDir, 0700)
	if err != nil {
		logger.Error.Fatal(err)
	}
	eventsChan := contrive.GenerateEventsChannel(*args.NumberOfGroups)
	batch.WriteFiles(eventsChan, *args.Interval, *args.OutputDir, *args.BatchSize)

}
