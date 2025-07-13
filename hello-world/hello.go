package main

import "fmt"

// Separate concerns: "domain" code from the outside world (side-effects)
// Why? To make it easier to test

// Domain is the "Hello world!" string
func Hello() string {
	return "Hello, world!"
}

func main() {
	// Println (printing to stdout) is a side effect
	fmt.Println(Hello())
}
