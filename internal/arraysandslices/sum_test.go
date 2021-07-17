package arraysandslices_test

import (
	"reflect"
	"testing"

	"github.com/joshuaswickirl/learn-go-with-tests/internal/arraysandslices"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}

	got := arraysandslices.Sum(numbers)
	want := 6

	if got != want {
		t.Errorf("got %d, want %d, given %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := arraysandslices.SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v, got, want", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := arraysandslices.SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := arraysandslices.SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}
