package iteration_test

import (
	"fmt"
	"testing"

	"github.com/joshuaswickirl/learn-go-with-tests/internal/iteration"
)

func TestRepeat(t *testing.T) {
	repeated := iteration.Repeat("a")
	expected := "aaaaaa"
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iteration.Repeat("a")
	}
}

func ExampleRepeat() {
	repeated := iteration.Repeat("a")
	fmt.Println(repeated)
	// Output: aaaaaa
}
