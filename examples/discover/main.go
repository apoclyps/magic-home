package main

import (
	"fmt"
	"os"
	"time"

	magichome "github.com/apoclyps/magic-home/pkg"
)

// Discover a list of available devices on the local network.
func main() {
	fmt.Print("Discovering ")

	go func() {
		for {
			fmt.Print(".")
			time.Sleep(100 * time.Millisecond)
		}
	}()

	devices, err := magichome.Discover(magichome.DiscoverOptions{
		BroadcastAddr: magichome.DEFAULT_BROADCAST_ADDR,
		Timeout:       1,
	})
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Print("\n\nDiscovered the following devices:\n\n")

	fmt.Println("IP         \t| ID         \t| Model")
	fmt.Println("-----------------------------------")
	for _, device := range *devices {
		fmt.Printf("%s\t| %s\t| %s\n", device.IP, device.ID, device.Model)
	}
}
