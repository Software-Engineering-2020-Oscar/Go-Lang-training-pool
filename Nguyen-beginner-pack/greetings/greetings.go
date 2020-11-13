// First
// // Declare  a greetings package to collect related functions
// package greetings

// import (
// 	// import the Go standard library errors package
// 	"errors"
// 	"fmt"
// )

// // Hello returns a greeting for the named person (capital - exported name)
// // The function takes a "name" parameter whose type is string, and returns a message string
// func Hello(name string) (string,error) {
// 	// If no name was given, return an error with a message.
// 	if name == "" {
// 		return "", errors.New("empty name")
// 	}
// 	// Return a greeting that embeds the name in a message.
// 	// the := operator is a shortcur for declaring and initializing  a variable in one line.
// 	// var message string
// 	// message = fmt.Sprintf("Hi, %v. Welcome!", name)
//     message := fmt.Sprintf("Hi, %v. Welcome!", name)
//     return message, nil
// }
// Second
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
		return name, errors.New("empty name")
	}
	// Create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
