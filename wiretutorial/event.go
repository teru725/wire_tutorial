package wiretutorial

import (
	"errors"
	"fmt"
)

func NewEvent(g Greeter) (Event, error) {
	if g.Grumy {
		return Event{}, errors.New("could not create event: greeter is grumy")
	}
	return Event{Greeter: g}, nil
}

// event
type Event struct {
	Greeter Greeter
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
