package main

import (
	"fmt"

	"github.com/ShubhzDev/Learn-Golang/Day-2/greetings"
)

func main() {
	message := greetings.Hello("Tommy")
	fmt.Println(message)
}
