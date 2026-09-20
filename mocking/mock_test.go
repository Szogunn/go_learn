package mocking

import (
	"bytes"
	"testing"
)

func TestCountdow(t *testing.T) {
	t.Run("basic Countdown", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{0}

		Countdown(buffer, spySleeper)

		got := buffer.String()
			want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}

		if spySleeper.Calls != 3 {
			t.Errorf("not enough calls to sleeper, want 3 but got %d", spySleeper.Calls)
		}

	})

}