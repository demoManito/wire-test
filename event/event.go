package event

import (
	"errors"
	"fmt"
	"time"
)

// Message msg
type Message string

func NewMessage(m string) Message {
	return Message(m)
}

// Greeter dependency: Message
type Greeter struct {
	Message Message // <- adding a Message field
	Grumpy bool
}

func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
	return g.Message
}

// Event dependency: Greeter
type Event struct {
	Greeter Greeter // <- adding a Greeter field
}

func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
