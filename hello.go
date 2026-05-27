package main

import "fmt"

// Domain code - Hello() should be separated from side effects - printing to the stdout
// Why? To make it easier to test our code.
func Hello() string {
	return "Hello, world"
}

func main() {
	fmt.Println(Hello())
}
