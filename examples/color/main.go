package main

import (
	"fmt"
	"net"
	"os"
	"time"

	magichome "github.com/apoclyps/magic-home/pkg"
	"github.com/apoclyps/magic-home/pkg/lights"
)

// Cycle an array of color presets for a device by setting color to each color with a delay between changes.
func main() {
	// Create a new Magic Home LED Strip Controller
	d, err := magichome.NewDevice(net.ParseIP("192.168.0.50"), "", "", "")
	if err != nil {
		fmt.Printf("Error configuring device: %s\n", err)
		os.Exit(1)
	}

	// Turn LED Strip Device On
	_, err = d.Power(true)
	if err != nil {
		fmt.Printf("Error powering device: %s\n", err)
		os.Exit(1)
	}

	var colors []lights.Color = []lights.Color{
		lights.White(),
		lights.Black(),
		lights.Red(),
		lights.Blue(),
		lights.Green(),
		lights.Yellow(),
		lights.Pink(),
		lights.Cyan(),
		lights.Silver(),
		lights.Gray(),
		lights.Maroon(),
		lights.Olive(),
		lights.Purple(),
		lights.Teal(),
		lights.Navy(),
		lights.Orange(),
	}

	// Cycle device colors
	for _, l := range colors {
		err := d.SetDeviceColor(l)
		if err != nil {
			fmt.Println("Error setting device color: ", err)
			os.Exit(1)
		}

		time.Sleep(3 * time.Second)
	}

	// Turn LED Strip Device On
	_, err = d.Power(false)
	if err != nil {
		fmt.Printf("Error powering device: %s\n", err)
		os.Exit(1)
	}
}
