package cmd

import (
	"fmt"
	"net"

	magichome "github.com/apoclyps/magic-home/pkg"
	"github.com/apoclyps/magic-home/pkg/lights"
	"github.com/spf13/cobra"
)

// deviceCmd represents the device command
var deviceCmd = &cobra.Command{
	Use:   "device",
	Short: "Control device of specific devices",
	Long: `Control device of a single specific device or all devices.

	Switch between On/Off device for one or more devices by providing the desired device.
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ip, _ := cmd.Flags().GetString("ip")
		hex, _ := cmd.Flags().GetString("hex")
		color, _ := cmd.Flags().GetString("color")
		return device(ip, hex, color, args)
	},
}

// device will validate the IP address for a single device (if provided) otherwise
// defaults to using all devices available on the local network.
// The color is set for one or more device's using either an exact match for
// a hex value or the name of a color preset.
func device(ip string, hex string, name string, args []string) error {
	var color lights.Color

	if hex != "" {
		cv, err := lights.NewValue(hex)
		if err != nil {
			return err
		}
		color, err = cv.GetColorByHex()
		if err != nil {
			return err
		}
	} else if name != "" {
		cv, err := lights.NewValue(hex)
		if err != nil {
			return err
		}
		color, err = cv.GetColorByName()
		if err != nil {
			return err
		}
	}

	if ip != "" && !magichome.IsPrivateIpv4(ip) {
		return fmt.Errorf("error while validating ip: %s", ip)
	} else if ip != "" {
		d, err := magichome.NewDevice(net.ParseIP(ip), "", "", "")
		if err != nil {
			return err
		}
		if err := d.SetDeviceColor(color); err != nil {
			return err
		}
	} else {
		devices, _ := magichome.Discover(magichome.DiscoverOptions{
			BroadcastAddr: magichome.DEFAULT_BROADCAST_ADDR,
			Timeout:       1,
		})

		for _, device := range *devices {
			if err := device.SetDeviceColor(color); err != nil {
				return err
			}
		}
		if len(*devices) == 0 {
			fmt.Printf("No devices to turn: on: \n")
		}
	}
	return nil
}

func init() {
	rootCmd.AddCommand(deviceCmd)
	deviceCmd.Flags().StringP("ip", "i", "", "Specify a single IP address to enable device for a specific devices")
	deviceCmd.Flags().StringP("hex", "x", "", "Specify a hex color to use for illumination")
	deviceCmd.Flags().StringP("color", "n", "", "Specify a color name to use for illumination")
}
