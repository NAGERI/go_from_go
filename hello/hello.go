package main

import (
	"fmt"

	"skylark.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Moiram")
	fmt.Println(message)
}
