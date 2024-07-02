package main

import (
	"fmt"

	"github.com/ShubhzDev/Learn-Golang/GoModule-2/greetings"
)

func main() {
	message := greetings.Hello("Tommy")
	fmt.Println(message)
}
