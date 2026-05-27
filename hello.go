package main

import "fmt"

const (
	spanish              = "Spanish"
	french               = "French"
	slovenian            = "Slovenian"
	englishHelloPrefix   = "Hello, "
	spanishHelloPrefix   = "Hola, "
	frenchHelloPrefix    = "Bonjour, "
	slovenianHelloPrefix = "Bogdej, "
)

// Domain code - Hello() should be separated from side effects - printing to the stdout
// Why? To make it easier to test our code.
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case slovenian:
		prefix = slovenianHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Joža", "Slovenian"))
}
