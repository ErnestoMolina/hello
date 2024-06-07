package main

import (
	"fmt"
	"log"

	"github.com/ErnestoMolina/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	fmt.Println()
	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
