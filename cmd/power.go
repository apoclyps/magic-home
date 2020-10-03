package cmd

import (
	"fmt"
	"net"

	magichome "github.com/apoclyps/magic-home/pkg"
	"github.com/spf13/cobra"
)

// powerCmd represents the power command
var powerCmd = &cobra.Command{
	Use:   "power",
	Short: "Control power of specific devices",
	Long: `Control power of a single specific device or all devices.

	Switch between On/Off power for one or more devices by providing the desired power.
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ip, _ := cmd.Flags().GetString("ip")

		var state = "off"
		power, _ := cmd.Flags().GetBool("power")
		if power {
			state = "on"
		}

		if ip != "" {
			fmt.Printf("Setting power to `on` for %s\n", ip)
			d, err := magichome.NewDevice(net.ParseIP(ip), "", "", "")
			if err != nil {
				fmt.Printf("Error configuring device: %s\n", err)
				return err
			}
			_, err = d.Power(power)
			if err != nil {
				fmt.Printf("Error powering device: %s\n", err)
				return err
			}
		} else {
			devices, _ := magichome.Discover(magichome.DiscoverOptions{
				BroadcastAddr: magichome.DEFAULT_BROADCAST_ADDR,
				Timeout:       1,
			})

			for _, device := range *devices {
				fmt.Printf("Setting power to %s for %s\n", state, device.IP.String())
				_, err := device.Power(power)
				if err != nil {
					fmt.Printf("Error powering device: %s\n", err)
					return err
				}
			}
			if len(*devices) == 0 {
				fmt.Printf("No devices to power: %s\n", state)
			}

		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(powerCmd)
	powerCmd.Flags().StringP("ip", "i", "", "Specify a single IP address to enable power for a specific devices")
	powerCmd.Flags().BoolP("power", "s", false, "Provide a boolean flag to enable (true) or disable (false) power ")
}
