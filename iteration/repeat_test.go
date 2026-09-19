package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected %q got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 6)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b", 5)
	fmt.Println(repeated)
	// Output: bbbbb
}
