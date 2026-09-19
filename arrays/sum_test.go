package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAllTailes(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTailes([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTailes([]int{}, []int{2, 3, 1})
		want := []int{0, 4}

		checkSums(t, got, want)
	})

}

func BenchmarkSumAll(b *testing.B) {
	for b.Loop() {
		SumAllTailes([]int{}, []int{0, 9})
	}
}
