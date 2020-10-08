package lights_test

import (
	"testing"

	"github.com/apoclyps/magic-home/pkg/lights"
)

func TestHexToColor(t *testing.T) {
	hex := "0000FF"
	actual, err := lights.HexToColor(hex)
	expected := lights.Color{R: 0, G: 0, B: 255, W: 255}

	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("actual %d expected %d", actual, expected)
	}
}

func TestHexToColors(t *testing.T) {
	colorTests := []struct {
		name  string
		hex   string
		color lights.Color
	}{
		{name: "Red", hex: "FF0000", color: lights.Color{R: 255, G: 0, B: 0, W: 255}},
		{name: "Green", hex: "00FF00", color: lights.Color{R: 0, G: 255, B: 0, W: 255}},
		{name: "Blue", hex: "0000FF", color: lights.Color{R: 0, G: 0, B: 255, W: 255}},
		{name: "White", hex: "FFFFFF", color: lights.Color{R: 255, G: 255, B: 255, W: 255}},
		{name: "Black", hex: "000000", color: lights.Color{R: 0, G: 0, B: 0, W: 255}},
	}

	for _, tt := range colorTests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := lights.HexToColor(tt.hex)
			if err != nil {
				t.Errorf("%#v got %d", tt.name, err)
			}
			if actual != tt.color {
				t.Errorf("%#v actual %d expected %d", tt.name, actual, tt.color)
			}
		})

	}
}
