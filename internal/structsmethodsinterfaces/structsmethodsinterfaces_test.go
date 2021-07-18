package structsmethodsinterfaces_test

import (
	"testing"

	"github.com/joshuaswickirl/learn-go-with-tests/internal/structsmethodsinterfaces"
)

func TestPerimeter(t *testing.T) {
	r := structsmethodsinterfaces.Rectangle{10.0, 10.0}
	got := r.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

type Shape interface {
	Area() float64
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{
			name:    "rectangle",
			shape:   structsmethodsinterfaces.Rectangle{Width: 10.0, Height: 10.0},
			hasArea: 100.0,
		},
		{
			name:    "circle",
			shape:   structsmethodsinterfaces.Circle{Radius: 10.0},
			hasArea: 314.1592653589793,
		},
		{
			name:    "triangle",
			shape:   structsmethodsinterfaces.Triangle{Base: 12, Height: 6},
			hasArea: 36.0,
		},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v, got %g, want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
