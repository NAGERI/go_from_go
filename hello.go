package main

import (
	"fmt"

	"skylark.com/greetings"
)

const englishHelloPrefix string = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World!"
	}
	return englishHelloPrefix + name
}
func main() {
	fmt.Println(Hello("Cedric"))

	// Get a greeting message and print it.
	message := greetings.Hello("Moira")
	fmt.Println(message)
}
