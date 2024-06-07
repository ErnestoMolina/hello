package main

import (
	"fmt"

	"github.com/ErnestoMolina/greetings"
)

func main() {
	fmt.Println()
	message := greetings.Hello("Helverzon")
	fmt.Println(message)
}
