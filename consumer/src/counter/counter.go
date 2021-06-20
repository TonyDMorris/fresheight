package counter

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/tonydmorris/fresh8-consumer/src/logger"
)

type Counter struct {
	viewed       int
	interacted   int
	clickThrough int
}

func New() *Counter {
	return &Counter{}
}
//CountEvents reads the given dir and processes the event information then deletes the file , 
//if no event files are found the process sleeps for 1 second 
func (c *Counter) CountEvents(dir string) {
	files, err := ioutil.ReadDir(dir)
	// error reading the files we should exit
	if err != nil {
		logger.Error.Fatal(err)
	}
	// if there are no files we sleep for 1 second to rate limit the operation and save cpu cycles
	if len(files) == 0 {
		time.Sleep(time.Second * 1)
	}
	for _, f := range files {
		c.countFile(fmt.Sprintf("%v/%v", dir, f.Name()))
	}
}
//countFile takes the lines from the file and counts the different types 
func (c *Counter) countFile(fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		// unable to open file , log out reason
		logger.Error.Print(err)
		return
	}
	scanner := bufio.NewScanner(file)
// for the purose of this application we unmarshal into an interface and count the types as we don't use any other information 
	for scanner.Scan() {
		var event map[string]interface{}
		err := json.Unmarshal(scanner.Bytes(), &event)
		if err != nil {
			logger.Error.Printf("unable to unmarshal event with ERROR: %v", err)
		}

		switch eventType := event["type"]; eventType {
		case "Viewed":
			c.viewed++
		case "Click-Through":
			c.clickThrough++
		case "Interacted":
			c.interacted++
		default:
			logger.Error.Print("got an unknown type when counting")
		}
	}
	if err := scanner.Err(); err != nil {
		//error reading lines
		logger.Error.Printf("error reading file lines  of file with filename %v and error of %v", fileName, err)
	}
	err = file.Close()
	if err != nil {
		logger.Error.Printf("error closing file with filename %v and error of %v", fileName, err)
	}
	err = os.Remove(fileName)
	if err != nil {
		// if unable to delete file counts will be innaccurate so exit program
		logger.Error.Fatal("unable to delete file shutting down")
	}

}
// LogCount formats the count structs values into a presented string 
func (c *Counter) LogCount() {
	fmt.Printf("\"Viewed\": %v \n \"Interacted\": %v \n \"Click-Through\": %v \n", c.viewed, c.interacted, c.clickThrough)
}
