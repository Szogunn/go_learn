package dependencyinjection

import (
	"bytes"
	"os"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("Greet with DI", func(t *testing.T) {
		buffer := bytes.Buffer{}
		pointerToBuffer := &buffer
		Greet(pointerToBuffer, "Chris")

		got := buffer.String()
		want := "Hello, Chris\n"

		if got != want {
			t.Errorf("got %s and want %s", got, want)
		}
	})

	t.Run("Greet with default Writer", func(t *testing.T) {
		Greet(os.Stdout, "Mateusz")
	})
}
