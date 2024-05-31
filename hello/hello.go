package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	// Use the log library and set the flags to disable time, source file and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Henry")

	if err != nil {
		log.Fatal(err)
	}

	// if no error was retuned , print the returned message
	fmt.Println(message)

	// A slice of names
	names := []string{"Henry", "Mary", "John", "Gladys"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	// if no error, print the map to the console
	fmt.Println(messages)
}
