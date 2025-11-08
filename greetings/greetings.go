package greetings

import "fmt"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("Hello, %s!", name)
}

func Goodbye(name string) string {
	if name == "" {
		name = "Friend"
	}
	return fmt.Sprintf("Goodbye, %s!", name)
}
