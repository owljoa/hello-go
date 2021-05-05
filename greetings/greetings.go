package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// if no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// if a name was received, return a value that embeds the name
	// in a greeting messagae
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
