package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(3)

	names := []string{"Max", "Joe", "Angela"}


	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
