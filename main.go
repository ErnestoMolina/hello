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

	names := []string{"Gabriela", "Ernesto", "Leider"}

	message, err := greetings.Hellos(names)
	// message, err := greetings.Hello("Juanito")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
