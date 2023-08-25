package wiretutorial

import "time"

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
