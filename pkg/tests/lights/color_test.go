package lights_test

import (
	"testing"

	"github.com/apoclyps/magic-home/pkg/lights"
)

func TestColor(t *testing.T) {
	actual, err := lights.NewColor(255, 255, 255, 255)
	expected := lights.Color{R: 255, G: 255, B: 255, W: 255}

	if err != nil {
		t.Error(err)
	}

	if actual != expected {
		t.Errorf("actual %d expected %d", actual, expected)
	}
}

func TestColors(t *testing.T) {
	colorTests := []struct {
		name  string
		r     uint8
		g     uint8
		b     uint8
		w     uint8
		color lights.Color
		err   error
	}{
		{name: "Red", r: 255, g: 0, b: 0, w: 255, color: lights.Color{R: 255, G: 0, B: 0, W: 255}, err: nil},
		{name: "Green", r: 0, g: 255, b: 0, w: 255, color: lights.Color{R: 0, G: 255, B: 0, W: 255}, err: nil},
		{name: "Blue", r: 0, g: 0, b: 255, w: 255, color: lights.Color{R: 0, G: 0, B: 255, W: 255}, err: nil},
		{name: "White", r: 255, g: 255, b: 255, w: 255, color: lights.Color{R: 255, G: 255, B: 255, W: 255}, err: nil},
		{name: "Black", r: 0, g: 0, b: 0, w: 0, color: lights.Color{R: 0, G: 0, B: 0, W: 0}, err: nil},
	}

	for _, tt := range colorTests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := lights.NewColor(tt.r, tt.g, tt.b, tt.w)
			if err != nil {
				t.Errorf("%#v actual %d expected %d", tt.name, err, tt.err)
			}
			if actual != tt.color {
				t.Errorf("%#v actual %d expected %d", tt.name, actual, tt.color)
			}
		})

	}
}
