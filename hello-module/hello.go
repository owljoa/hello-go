package main

import (
	"fmt"
	"log"

	"github.com/owljoa/hello-go/greetings"
)

func main() {
	// Set properties of the predefined Logger. including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Owljoa", "JoaOwl", "OlaJow"}

	// Request a greeting message for the names
	message, err := greetings.HelloAll(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// message to the console.
	fmt.Println(message)
}
