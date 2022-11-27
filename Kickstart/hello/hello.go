package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)


func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	
	// A slice of names
	names := []string{"Chris", "Chelsea", "Yogayee"}

	// Request greeting messages for names
	messages, err := greetings.Hellos(names)
	if err != nil {
		// This will exit program
		log.Fatal(err)
	}

	fmt.Println(messages)
}


