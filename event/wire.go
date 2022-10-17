//+build wireinject

package event

import (
	"github.com/google/wire"
)

func Init(msg string) (Event, error) {
	wire.Build(NewMessage, NewGreeter, NewEvent)
	return Event{}, nil
}