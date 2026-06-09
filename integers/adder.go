// Package integers, explains how integers work in Go (as you would expect).
// Adding a comment like this add the top of the file (or in doc.go for long
// comments) serves as a documentation, that you can then use with `go doc package`
// or pkgsite. If you don't have a comment like this, then golangci-lint displays a
// warning and nudges you to follow Go's documentation best practices.
package integers

// Add takes two integers and returns the sum of them.
func Add(x, y int) int {
	return x + y
}
