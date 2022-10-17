package main

import (
	"fmt"
	"os"

	"wire-test/event"
)

func main() {
	event,err := event.Init("Hi there! 1")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	event.Start()
}