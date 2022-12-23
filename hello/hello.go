package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Eymard")
	fmt.Println(message)

}
