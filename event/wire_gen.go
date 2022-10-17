// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package event

// Injectors from wire.go:

func Init(msg string) (Event, error) {
	message := NewMessage(msg)
	greeter := NewGreeter(message)
	event, err := NewEvent(greeter)
	if err != nil {
		return Event{}, err
	}
	return event, nil
}