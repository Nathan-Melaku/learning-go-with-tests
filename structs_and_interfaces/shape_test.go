package structsandinterfaces

import (
	"testing"
)

func checksum(got, want float64, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	checksum(got, want, t)
}

func TestArea(t *testing.T) {
	t.Run("shapes area", func(t *testing.T) {
		areaTests := []struct {
			name    string
			shape   Shape
			hasArea float64
		}{
			{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
			{name: "Cirlcle", shape: Circle{10}, hasArea: 314.1592653589793},
			{name: "Trianlgle", shape: Triangle{12, 6}, hasArea: 36},
		}

		for _, tt := range areaTests {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v %g wnat %g", tt.shape, got, tt.hasArea)
			}
		}
	})
}
