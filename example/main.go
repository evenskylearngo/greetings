package main

import (
	"fmt"

	"github.com/evenskylearngo/greetings/greetings"
)

func main() {
	fmt.Println(greetings.Hello("Andy"))
	fmt.Println(greetings.Goodbye("Andy"))
}
