package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("No name was given")
	}

	// If a name is received, return a value that embeds the name in a greeting message
	var message string = fmt.Sprintf(randomFormat(), name)
	
	return message, nil
}

// Init sets initial values for variables used in  the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greetings messages. The returned message is selected at random.
func randomFormat() string {
	// A slice of message strings
	formats :=  []string {
		"Hi %v. Welcome",
		"Greetings to see you, %v!",
		"Hail, %v! Well met!",
		"Hi %v from the future. This is past you from writing this code!",
	}
    // Return a randomly selected message format by specifying a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}