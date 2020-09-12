package main

import (
	"fmt"
	"net"
	"time"

	magichome "github.com/apoclyps/magic-home/pkg"
	"github.com/apoclyps/magic-home/pkg/lights"
)

// Cycle an array of color temperatures for a device by setting color to each color with a delay between changes.
func main() {
	// Create a new Magic Home LED Strip Controller
	controller, err := magichome.NewController(net.ParseIP("192.168.0.50"), 5577)
	if err != nil {
		fmt.Println("Magic Home Controller Error: ", err)
		return
	}

	// Turn LED Strip Controller On
	err = controller.SetState(magichome.On)
	if err != nil {
		fmt.Println("Set State Error: ", err)
		return
	}

	var temperatures = []lights.Color{
		lights.Candle(),
		lights.Tungsen40W(),
		lights.Tungsen100W(),
		lights.Halogen(),
		lights.CarbonArc(),
		lights.HighNoonSun(),
		lights.DirectSunlight(),
		lights.OvercastSky(),
		lights.ClearBlueSky(),
	}

	// Cycle device tempatures
	for _, l := range temperatures {
		err := controller.SetColor(l)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		time.Sleep(3 * time.Second)
	}

	// And finaly close the connection to LED Strip Controller
	err = controller.Close()
	if err != nil {
		fmt.Println("Magic Home Controller Error: ", err)
		return
	}
}
