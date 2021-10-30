### Magic Home

![](https://banners.beyondco.de/Magic%20Home.png?theme=light&packageManager=&packageName=go+install+github.com%2Fapoclyps%2Fmagic-home%40latest&pattern=fourPointStars&style=style_2&description=Controlling+your+Magic+Home+LED+light+strips.&md=1&showWatermark=0&fontSize=150px&images=light-bulb)

Magic Home is a CLI & library for controlling RGB LED light strips.

### CLI Usage

```bash
> go build .

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

### Usage

Discover devices on the network.

```bash
> ./magic-home discover

Discovering ..........

Discovered the following devices:

IP         	| ID         	| Model
-----------------------------------
192.168.0.30	| C44F33AEB69A	| AK001-ZJ210
192.168.0.32	| D8F15B8446DF	| AK001-ZJ210
```

Power on a device on the network (using the previous color set).

```bash
> ./magic-home power --ip=192.168.0.30 --power=true

Setting power to `on` for 192.168.0.30
```

Set the color using a preset name for a target device on the network.

```bash
> ./magic-home device --ip=192.168.0.30 --color=blue

Setting device color {0 0 255 0} for 192.168.0.30
```

Set the color using a hex value for a target device on the network.

```bash
> ./magic-home device --ip=192.168.0.30 --hex=0000ff

Setting device color {0 0 255 0} for 192.168.0.30
```

Play a scene for a target device on the network.

```bash
> ./magic-home scene --ip=192.168.0.30 --color=blue

Setting device color {0 0 255 0} for 192.168.0.30
Setting device color {0 0 55 0} for 192.168.0.30
Setting device color {0 0 255 0} for 192.168.0.30
Setting device color {0 0 55 0} for 192.168.0.30
Setting device color {0 0 255 0} for 192.168.0.30
Setting device color {0 0 55 0} for 192.168.0.30
Setting device color {0 0 255 0} for 192.168.0.30
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


### Disclaimer

> As this library controls physical hardware on a LAN, you should use your judgment to decide if the changes you are making may negatively impact your hardware.
