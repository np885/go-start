package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(3)



	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
