package contrive

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Type string    `json:"type"`
	Data EventData `json:"event_data"`
}

type EventData struct {
	ViewID        string    `json:"viewId"`
	EventDateTime time.Time `json:"eventDateTime"`
}

//generateEvent creates a single viewed event or either a pair of viewed and interacted or viewed and click-through or al three
// based on specified % rate
func generateEvent() []Event {
	// make a roll between 1 - 100
	roll := rand.Intn(100)
	// generate uuid
	uuid := generateUUID()
	// create mock records
	viewed := Event{
		Type: "Viewed",
		Data: EventData{
			ViewID:        uuid,
			EventDateTime: time.Time{},
		},
	}

	clickThrough := Event{
		Type: "Click-Through",
		Data: EventData{
			ViewID:        uuid,
			EventDateTime: time.Time{},
		},
	}

	interacted := Event{
		Type: "Interacted",
		Data: EventData{
			ViewID:        uuid,
			EventDateTime: time.Time{},
		},
	}
	// there is a better way to do this
	// return the events at the required rate
	if roll <= 5 {
		return []Event{clickThrough, viewed}
	} else if roll > 5 && roll <= 10 {
		return []Event{interacted, viewed}
	} else if roll > 10 && roll <= 15 {
		return []Event{clickThrough, interacted, viewed}
	} else {
		return []Event{viewed}
	}
}

func GenerateEventsChannel(numberOfGroups int) chan []Event {
	// create chan to receive events
	eventsChan := make(chan []Event)
	for i := 0; i < numberOfGroups; i++ {
		// spin up required number of event producers
		go func(c chan []Event) {
			// infinate loop of generating events with a fake rate limit
			for range time.Tick(time.Millisecond * time.Duration(100)) {
				eventsChan <- generateEvent()
			}
		}(eventsChan)
	}
	return eventsChan
}

func generateUUID() string {
	id := uuid.New()
	return id.String()
}
