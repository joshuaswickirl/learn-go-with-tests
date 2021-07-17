package integers_test

import (
	"fmt"
	"testing"

	"github.com/joshuaswickirl/learn-go-with-tests/internal/integers"
)

func TestAdder(t *testing.T) {
	sum := integers.Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := integers.Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
