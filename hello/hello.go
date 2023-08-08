package main

import (
	"fmt"
	"runtime"

	"skylark.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Moiram")
	fmt.Println(message)

	fmt.Printf("current system: %s/%s", runtime.GOOS, runtime.GOARCH)
}
