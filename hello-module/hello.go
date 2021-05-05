package main

import (
	"fmt"

	"github.com/owljoa/hello-go/greetings"
)

func main() {
	// Get a greetings message and print it.
	message := greetings.Hello("Owljoa")
	fmt.Println(message)
}
