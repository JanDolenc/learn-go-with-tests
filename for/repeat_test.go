package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// Only the body of the loop is timed
func BenchmarkRepeat(b *testing.B) {
	// ... setup ...
	for b.Loop() {
		// ... code to mesasure ...

		Repeat("a")
	}
	// ... clean up ...
}
