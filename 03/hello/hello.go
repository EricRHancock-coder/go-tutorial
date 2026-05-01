// Packge main print randomly generated message from package greetings
package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, sourc file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("Gladys")
	// if an error was returned, print it to the console and
	// exit the program
	if err != nil {
		log.Fatal(err)
	}

	// if no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}
