package main

import (
	"time"

	magichome "github.com/apoclyps/magic-home/pkg"
	"github.com/apoclyps/magic-home/pkg/lights"
)

// Cycle through an array of colors with a delay between changes for a number of iterations.
func main() {
	devices, _ := magichome.Discover(magichome.DiscoverOptions{
		BroadcastAddr: magichome.DEFAULT_BROADCAST_ADDR,
		Timeout:       3,
	})

	var colorArray = []lights.Color{
		lights.Blue(),
		lights.Red(),
		lights.Blue(),
	}
	var iterations uint8 = 3
	var delay time.Duration = 250 * time.Millisecond

	s := magichome.NewScene(devices, colorArray, iterations, delay)
	s.Play()
}
