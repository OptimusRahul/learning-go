package greetings

import "fmt"

// Hello returns a greeting for the named person
func Hello(name string) string {
	// Return a greeting that embeds the name in message

	// This is the long way of declaring and initializing the variable
	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)

	// this is the shorthand for declaring and initializing the variable
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
