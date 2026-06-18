package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// Only the body of the loop is timed
func BenchmarkRepeat(b *testing.B) {
	// ... setup ...
	for b.Loop() {
		// ... code to mesasure ...

		Repeat("a", 5)
	}
	// ... clean up ...
}

func ExampleRepeat() {
	repeated := Repeat("s", 7)

	fmt.Printf("Here is your character repeated: %s", repeated)
	// Output: Here is your character repeated: sssssss
}
