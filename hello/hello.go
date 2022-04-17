package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	//Get a greeting message and print it.
	// TODO: implement error checking
	message, _ := greetings.Hello("Bertrand")
	fmt.Println(message)
}
