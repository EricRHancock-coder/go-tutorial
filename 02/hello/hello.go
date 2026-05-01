package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable Printing
	// the time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a getting message.
	message, err := greetings.Hello("")
	// if an error was returned, print it to the console
	// and exit the program
	if err != nil {
		log.Fatal(err)
	}

	// if no error was returned, print the returned message
	// to th econsole
	fmt.Println(message)
}
