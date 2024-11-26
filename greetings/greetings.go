package greetings

import (
	"fmt"
	"errors")

//Funktionen mit Großen Anfangsbuchstabe sind public -> exported Name
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	//:= Declaring and init at the same time
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
