package magichome

import (
	"time"

	"github.com/apoclyps/magic-home/pkg/lights"
)

// A Scene represents a list of Device's that will iterate through
// an array of colors a specified amount of times with a delay between
// each color change.
type Scene struct {
	devices    *[]Device
	colors     []lights.Color
	iterations uint8
	delay      time.Duration
}

// NewScene initializes a Scene that can be played on one or more devices.
// Using an array of provided colors, each device will display the next colour
// for the specified iteration with a delay between colour changes.
func NewScene(devices *[]Device, colors []lights.Color, iterations uint8, delay time.Duration) Scene {
	scene := Scene{
		devices:    devices,
		colors:     colors,
		iterations: iterations,
		delay:      delay,
	}
	return scene
}

// Play iteratees through a list of colors on all devices with a delay
// betweeen each colour change.
func (s Scene) Play() {
	for _, c := range s.colors {
		for _, d := range *s.devices {
			d.SetDeviceColor(c)
			time.Sleep(s.delay)
		}
	}
}
