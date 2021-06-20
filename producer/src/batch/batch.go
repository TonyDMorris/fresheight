package batch

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/tonydmorris/fresh8-producer/src/contrive"
	"github.com/tonydmorris/fresh8-producer/src/logger"
)

type batchQueue struct {
	batches      [][]contrive.Event
	currentBatch []contrive.Event
	mu           sync.Mutex
	batchSize    int
	outputDir    string
}

const concurrency = 10

func WriteFiles(eventsChan chan []contrive.Event, interval int, dir string, batchSize int) {

	ba := batchQueue{batchSize: batchSize, outputDir: dir}
	// goroutine to intiate batch writing
	go func(ba *batchQueue) {

		ba.writeBatchAtInterval(time.Second * time.Duration(interval))
	}(&ba)
	// concurrently add events to batches
	for {
		wg := sync.WaitGroup{}
		wg.Add(concurrency)
		for i := 0; i < concurrency; i++ {
			// theoretically there would be some business logic here that we would want to execute concurrently
			go func(wg *sync.WaitGroup, eventsChan chan []contrive.Event, ba *batchQueue) {
				event := <-eventsChan
				ba.processAndWriteEvents(event)
				wg.Done()
			}(&wg, eventsChan, &ba)
		}
		wg.Wait()
	}
}

func (ba *batchQueue) writeBatch() {
	// format the filename
	fileName := fmt.Sprintf("%v/%v.json", ba.outputDir, time.Now().Format("Jan-02-15:04:05.000000000"))
	// create the file
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		logger.Error.Panic(err)
	}
	// make a new json encoder
	enc := json.NewEncoder(f)
	// append each json obj to the file
	for _, event := range ba.currentBatch {
		enc.Encode(event)
	}
	//clear the batch
	ba.currentBatch = ba.currentBatch[:0]

}

// processAndWriteEvents checks the current batch size against the max batch size and either continues or writes the batch to aa file
func (ba *batchQueue) processAndWriteEvents(e []contrive.Event) {
	// lock the batchQueue
	ba.mu.Lock()
	// get total size of combined
	totalSize := len(ba.currentBatch) + len(e)
	// if total == batch size add to current batch add batch to queue and clear current batch
	if totalSize == ba.batchSize {
		ba.currentBatch = append(ba.currentBatch, e...)
		ba.writeBatch()
		// if its less add events to current batch
	} else if totalSize < ba.batchSize {
		ba.currentBatch = append(ba.currentBatch, e...)
		// if total size > batchSize add batch to queue clear current batch and add ne events
	} else {
		ba.writeBatch()
		ba.currentBatch = append(ba.currentBatch, e...)
	}
	ba.mu.Unlock()
}

func (ba *batchQueue) writeBatchAtInterval(i time.Duration) {
	// range over time at the given rate and write batches out to fs weather the batch is full or not
	for range time.Tick(i) {
		ba.mu.Lock()
		ba.writeBatch()
		ba.mu.Unlock()
	}
}
