package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	return fmt.Sprintf("Hello, %s!", name), nil
}

func Goodbye(name string) string {
	if name == "" {
		name = "Friend"
	}
	return fmt.Sprintf("Goodbye, %s!", name)
}

func Plus(a, b int) int {
	return a + b
}
func Minus(a, b int) int {
	return a - b
}
