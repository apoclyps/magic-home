package main

import (
	"fmt"
	"net"
	"time"

	magichome "github.com/apoclyps/magic-home/pkg"
	"github.com/apoclyps/magic-home/pkg/lights"
)

// Blink device by setting color to white, and toggling device power on/off
func blink(controller *magichome.Controller) {
	time.Sleep(1 * time.Second)

	// Set color of LED Strip to white
	err := controller.SetColor(lights.Candle())
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	err = controller.SetState(magichome.On)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Sleep again a few seconds to avoid spam
	time.Sleep(1 * time.Second)

	// Turn LED Strip Controller off
	err = controller.SetState(magichome.Off)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}

func main() {
	// Create a new Magic Home LED Strip Controller
	controller, err := magichome.NewController(
		net.ParseIP("192.168.0.50"), 5577,
	)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	for i := 0; i < 3; i++ {
		blink(controller)
	}

	// And finaly close the connection to LED Strip Controller
	err = controller.Close()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}
