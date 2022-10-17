package main

import (
	"fmt"
	"os"
)

func main() {
	event,err := InitEvent("Hi wire!")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	event.Start()
}