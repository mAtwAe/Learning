package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
		// Return an error if the name is empty.
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randokmFormat(), name)
	return message, nil
}

func randokmFormat() string {
	formats := []string{
		"hi, %v. Welcome!",
		"Hello, %v. Glad to see you!",
		"Greetings, %v. How are you?",
		"Hey, %v. Nice to meet you!",
		"Welcome, %v. Hope you're doing well!",
	}

	return formats[rand.Intn(len(formats))]
}
