### Magic Home

Magic Home is a CLI & library for controlling RGB LED light strips.

### CLI Usage

```bash
> ./magic-home

Control Magic Home enabled devices.

  Allows control of LED strips and bulbs compatiable with the Magic Home App (Devices must be connected to the same LAN).

Usage:
  magic-home [command]

Available Commands:
  device      Control device of specific devices
  discover    Discover Magic Home Devices
  help        Help about any command
  power       Control power of specific devices
  scene       Control scene using specific devices

Flags:
  --config string       config file (default is $HOME/.magic-home.yaml)
  -h, --help            help for magic-home

Use "magic-home [command] --help" for more information about a command.
```

### Example

| Name                                           | Description                                                                                                    |
|---                                             |---                                                                                                             |
| [Blink](./examples/blink/main.go)              | Blink device by setting color to white, and toggling device power on/off.                                      |
| [Color](./examples/color/main.go)              | Cycle an array of color presets for a device by setting color to each color with a delay between changes.      |
| [Discover](./examples/discover/main.go)        | Discover a list of available devices on the local network.                                                     |
| [Scene](./examples/scene/main.go)              | Cycle through an array of colors with a delay between changes for a number of iterations.                      |
| [State](./examples/state/main.go)              | Power a device's on and then power the same device off.                                                        |
| [Temperature](./examples/temperature/main.go)  | Cycle an array of color temperatures for a device by setting color to each color with a delay between changes. |
