// Package iteration, includes examples how to do stuff repeatedly in Go.
package iteration

import "strings"

// Repeat returns character repeated n times.
func Repeat(character string, repeatCount int) string {
	var repeated strings.Builder
	for range repeatCount {
		repeated.WriteString(character)
	}
	return repeated.String()
}
