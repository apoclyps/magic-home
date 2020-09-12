package main

import (
	"fmt"
	"net"
	"time"

	magichome "github.com/apoclyps/magic-home/pkg"
	"github.com/apoclyps/magic-home/pkg/lights"
)

// Cycle an array of color presets for a device by setting color to each color with a delay between changes.
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
		err := controller.SetColor(l)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		time.Sleep(3 * time.Second)
	}

	// Turn LED Strip Controller off
	err = controller.SetState(magichome.Off)
	if err != nil {
		fmt.Println("Set State Error: ", err)
		return
	}

	// And finaly close the connection to LED Strip Controller
	err = controller.Close()
	if err != nil {
		fmt.Println("Magic Home Controller Error: ", err)
		return
	}
}
