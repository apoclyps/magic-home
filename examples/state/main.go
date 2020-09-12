package main

import (
	"fmt"
	"net"
	"os"

	magichome "github.com/apoclyps/magic-home/pkg"
)

// Power a device's on and then power the same device off.
func main() {
	ip := "192.168.0.50"

	d, err := magichome.NewDevice(net.ParseIP(string(ip)), "", "", "")
	if err != nil {
		fmt.Printf("Error configuring device: %s\n", err)
		os.Exit(1)
	}
	_, err = d.Power(true)
	if err != nil {
		fmt.Printf("Error powering device: %s\n", err)
		os.Exit(1)
	}

	_, err = d.Power(false)
	if err != nil {
		fmt.Printf("Error powering device: %s\n", err)
		os.Exit(1)
	}
}
