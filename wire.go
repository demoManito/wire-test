//+build wireinject

package main

import (
	"wire-test/event"

	"github.com/google/wire"
)

func InitEvent(msg string) (event.Event, error) {
	wire.Build(event.NewMessage, event.NewGreeter, event.NewEvent)
	return event.Event{}, nil
}