package greetings

import (
	"fmt"
	"errors"
)

// FUNCTION WITH ARGUMENTS
// In Go, a function whose name starts with a capital letter can be called by a function not in the same package
func Hello(name string) (string, error) {

	// IF NO NAME WAS GIVEN:
	if name == "" {
		return "", errors.New("empty name")
	}
	
	// the := operator is used as a shortcut for declaring and initializing a variable simultaneously
	// Sprintf = Substitute + Printf => substitutes variabels and then prints the string.
	message := fmt.Sprintf("Hi, %v, Welcome!", name)
	return message, nil
}
