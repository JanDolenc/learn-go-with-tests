package main

import "fmt"

// Separate concerns: "domain" code from the outside world (side-effects)
// Why? To make it easier to test

// Go has constants
// Creating constants (or variables) to capture meaning of values and sometimes improve performance
const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

// Domain is the "Hello world!" string
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return spanishHelloPrefix + name + "!"
	}

	return englishHelloPrefix + name + "!"
}

func main() {
	// Println (printing to stdout) is a side effect
	fmt.Println(Hello("Jan", ""))
}
