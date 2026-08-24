package main
import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main() {

	// set predefined logger's settings:
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	
	// get greetings function and print
	message, err := greetings.Hello("")

	// if an error was returned:
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(message)
}