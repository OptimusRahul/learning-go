package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// if no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in message

	// This is the long way of declaring and initializing the variable
	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)

	// this is the shorthand for declaring and initializing the variable
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
