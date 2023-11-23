package main

import (
	"fmt"
	"log"

	"example.com/greetings"

	"rsc.io/quote"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)

	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	message, err := greetings.Hello("Gladys")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
