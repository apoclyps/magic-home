package magichome

import (
	"fmt"
	"net"
	"os"

	"github.com/apoclyps/magic-home/pkg/lights"
)

// State represents the status of the device (on/off)
type State uint8

const (
	// On represents the switched on power for a device
	On State = 1
	// Off represents the switched off power for a device
	Off State = 0
)

// Device represents a Magic Home device (responds to a UDP broadcast)
type Device struct {
	IP    net.IP
	ID    string
	Model string
}

// New initializes a new Magic Home LED Device.
// A device comprises an IP to target the device on the network,
// an unique ID registered to the device, and the model of device
// responding a broadcast request.
func NewDevice(ip net.IP, id, string, model string) (*Device, error) {
	d := Device{
		IP:    ip,
		ID:    id,
		Model: model,
	}
	return &d, nil
}

// Power initialises a new device controller and send's a power state
// to the device to activate or deactive the device.
func (d *Device) Power(power bool) (*Controller, error) {
	var status State
	if power {
		status = On
	} else {
		status = Off
	}

	controller, err := NewController(d.IP, 5577)
	if err != nil {
		return nil, err
	}

	//defer controller.Close()
	err = controller.SetState(status)
	if err != nil {
		return controller, err
	}

	return controller, nil
}

// SetDeviceColor accepts a color, activates the power on that device
// to ensure it's active, and then set's the device colour. Once a device
// colour is set, it will remain set until changed or the power is deactivated.
func (d *Device) SetDeviceColor(c lights.Color) {
	fmt.Printf("Setting device to %t for %s\n", true, d.IP)

	controller, err := d.Power(true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = controller.SetColor(c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	controller.Close()
}
