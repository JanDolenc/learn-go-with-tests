package main

import "fmt"

// Separate concerns: "domain" code from the outside world (side-effects)
// Why? To make it easier to test

// Go has constants
// Creating constants (or variables) to capture meaning of values and sometimes improve performance
const englishHelloPrefix = "Hello, "

// Domain is the "Hello world!" string
func Hello(name string) string {
	return englishHelloPrefix + name + "!"
}

func main() {
	// Println (printing to stdout) is a side effect
	fmt.Println(Hello("world"))
}
