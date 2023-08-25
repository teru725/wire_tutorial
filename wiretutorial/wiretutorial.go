package wiretutorial

import (
	"errors"
	"fmt"
	"time"
)

type Message string

func NewMessage(phrase string) Message {
	return Message(phrase)
}

// greeter
type Greeter struct {
	Message Message
	Grumy   bool
}

func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumy: grumpy}
}

func (g Greeter) Greet() Message {

	return g.Message
}

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
