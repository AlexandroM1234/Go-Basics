package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
    var message string = (greetings.Hello("Alex"))
	
	fmt.Println(message)
}
