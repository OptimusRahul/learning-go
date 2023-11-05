package main

import (
	"fmt"
	"greetings"
	"log"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file and line number.
	log.SetPrefix("greeting: ")
	log.SetFlags(0)

	// A slice of names
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request a greeting message
	// Get a greeting message and print it.
	message, err := greetings.Hellos(names)
	// if an error was returne, print it to the console and
	// exit the program
	if err != nil {
		log.Fatal(err)
	}

	// if no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}
