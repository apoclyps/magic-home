package magichome

import (
	"fmt"
	"net"

	"github.com/apoclyps/magic-home/pkg/lights"
)

// Controller represents a Magic Home LED Strip Controller
type Controller struct {
	ip      net.IP
	port    uint16
	conn    net.Conn
	command Command
}

// New initializes a new Magic Home LED Strip Controller
func NewController(ip net.IP, port uint16) (*Controller, error) {
	c := Controller{
		ip:   ip,
		port: port,
		command: Command{
			on:    []byte{0x71, 0x23, 0x94},
			off:   []byte{0x71, 0x24, 0x95},
			color: []byte{0x31, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
	}

	var addr string = fmt.Sprintf("%s:%d", c.ip, c.port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	c.conn = conn

	return &c, nil
}

// SetState can be used to switch the LED Strip on (magichome.On) or off (magichome.Off)
func (c *Controller) SetState(s State) error {
	if s == On {
		_, err := c.conn.Write(c.command.on)
		return err
	}
	_, err := c.conn.Write(c.command.off)
	return err
}

// SetColor can be used to change the color of the LED Strip
func (c *Controller) SetColor(color lights.Color) error {
	bytesToSend := c.command.color

	bytesToSend[1] = byte(color.R)
	bytesToSend[2] = byte(color.G)
	bytesToSend[3] = byte(color.B)
	bytesToSend[4] = byte(color.W)

	var sum int
	for _, value := range bytesToSend {
		sum += int(value)
	}

	bytesToSend = append(bytesToSend, byte(sum&0xff))

	_, err := c.conn.Write(bytesToSend)
	if err != nil {
		return err
	}

	return nil
}

// Close closes the tcp connection to the LED Strip
func (c *Controller) Close() error {
	return c.conn.Close()
}
