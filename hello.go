package main

import "fmt"

// Domain code - Hello() should be separated from side effects - printing to the stdout
// Why? To make it easier to test our code.
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("Manca"))
}
