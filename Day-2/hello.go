package main

import (
	"fmt"

	"log"

	"github.com/ShubhzDev/Learn-Golang/Day-2/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Tommy")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{"Ram", "Sita", "Lakhsman"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
