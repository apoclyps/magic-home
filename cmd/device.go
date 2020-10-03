package cmd

import (
	"fmt"
	"net"
	"os"

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
	Run: func(cmd *cobra.Command, args []string) {
		ip, _ := cmd.Flags().GetString("ip")
		hex, _ := cmd.Flags().GetString("hex")
		color, _ := cmd.Flags().GetString("color")
		device(ip, hex, color, args)
	},
}

// device function sets color based on hex and name parameters while validating ip
func device(ip string, hex string, name string, args []string) {
	color := lights.GetColor(hex, name)

	if ip != "" && !magichome.IsPrivateIpv4(ip) {
		fmt.Printf("Error while validating ip: %s", ip)
		os.Exit(1)
	} else if ip != "" {
		d, err := magichome.NewDevice(net.ParseIP(ip), "", "", "")
		if err != nil {
			fmt.Println(err)
			return
		}
		d.SetDeviceColor(color)
	} else {
		devices, _ := magichome.Discover(magichome.DiscoverOptions{
			BroadcastAddr: magichome.DEFAULT_BROADCAST_ADDR,
			Timeout:       1,
		})

		for _, device := range *devices {
			device.SetDeviceColor(color)
		}
		if len(*devices) == 0 {
			fmt.Printf("No devices to turn: on: \n")
		}
	}
}

func init() {
	rootCmd.AddCommand(deviceCmd)
	deviceCmd.Flags().StringP("ip", "i", "", "Specify a single IP address to enable device for a specific devices")
	deviceCmd.Flags().StringP("hex", "x", "", "Specify a hex color to use for illumination")
	deviceCmd.Flags().StringP("color", "n", "", "Specify a color name to use for illumination")
}
