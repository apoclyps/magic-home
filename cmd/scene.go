package cmd

import (
	"fmt"
	"net"
	"time"

	magichome "github.com/apoclyps/magic-home/pkg"
	"github.com/apoclyps/magic-home/pkg/lights"
	"github.com/spf13/cobra"
)

// sceneCmd represents the scene command
var sceneCmd = &cobra.Command{
	Use:   "scene",
	Short: "Control scene using specific devices",
	Long: `Control scene using specific devices.

	Switch between On/Off scene for one or more scenes by providing the desired scene.
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ip, _ := cmd.Flags().GetString("ip")
		color, _ := cmd.Flags().GetString("color")
		return scene(ip, color, args)
	},
}

func createColorArray(c lights.Color, iterations int) []lights.Color {
	var arr []lights.Color

	var r uint8 = c.R
	if c.R == 255 {
		r = c.R - 200
	}
	var g uint8 = c.G
	if c.G == 255 {
		g = c.G - 200
	}
	var b uint8 = c.B
	if c.B == 255 {
		b = c.B - 200
	}

	modifiedColor := lights.Color{R: r, G: g, B: b, W: c.W}
	for i := 0; i < iterations; i++ {
		arr = append(arr, c)
		arr = append(arr, modifiedColor)
	}
	arr = append(arr, c)

	return arr
}

func scene(ip string, colorName string, args []string) error {
	c, err := lights.GetColorByName(colorName)
	if err != nil {
		return err
	}

	colorArray := createColorArray(c, 3)

	if ip != "" && !magichome.IsPrivateIpv4(ip) {
		fmt.Printf("Error while validating ip: %s", ip)
		return fmt.Errorf("error while validating ip: %s", ip)
	} else if ip != "" {
		for _, color := range colorArray {
			d, err := magichome.NewDevice(net.ParseIP(ip), "", "", "")
			if err != nil {
				fmt.Println(err)
				return err
			}
			d.SetDeviceColor(color)
			time.Sleep(100 * time.Millisecond)
		}
	} else {
		devices, _ := magichome.Discover(magichome.DiscoverOptions{
			BroadcastAddr: magichome.DEFAULT_BROADCAST_ADDR,
			Timeout:       3,
		})

		var iterations uint8 = 3
		var delay time.Duration = 250 * time.Millisecond

		s := magichome.NewScene(devices, colorArray, iterations, delay)
		s.Play()

		if len(*devices) == 0 {
			fmt.Printf("No scenes to turn: on: \n")
		}
	}
	return nil
}

func init() {
	rootCmd.AddCommand(sceneCmd)
	sceneCmd.Flags().StringP("ip", "i", "", "Specify a single IP address to enable scene for a specific scenes")
	sceneCmd.Flags().StringP("color", "n", "", "Specify a color name to use for illumination")
}
