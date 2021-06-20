package main

import (
	"os"
	"time"

	"github.com/tonydmorris/fresh8-consumer/src/counter"

	"github.com/tonydmorris/fresh8-consumer/src/args"
	"github.com/tonydmorris/fresh8-consumer/src/logger"
)

func main() {
	// init program args
	args.Init()
	// check input dir exists
	if _, err := os.Stat(*args.InputDir); os.IsNotExist(err) {
		logger.Error.Fatal(err)
	}
	// make new counter
	c := counter.New()
	// log out the current count every 5 seconds
	go func(c *counter.Counter) {
		for range time.Tick(time.Second * 5) {
			c.LogCount()
		}
	}(c)
	// count the events in the input dir forever
	for {
		c.CountEvents(*args.InputDir)
	}
}
