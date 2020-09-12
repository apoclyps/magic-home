package magichome

import (
	"fmt"
	"net"
	"strings"
	"time"
)

//DiscoverOptions stores configuration to find Device's on the local area network
type DiscoverOptions struct {
	BroadcastAddr string
	Timeout       uint8
}

// Send UDP packets to Device's on the local area network.
func send(conn *net.UDPConn, broadcastAddr *net.UDPAddr) {
	const commandPassword = "HF-A11ASSISTHREAD"

	_, err := conn.WriteToUDP([]byte(commandPassword), broadcastAddr)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

// Receives a list of Device's on the local area network.
func receive(conn *net.UDPConn, receivedDevice chan<- *Device) {
	for {
		buf := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			break
		}

		data := strings.Split(string(buf[:n]), ",")
		if len(data) == 3 {
			receivedDevice <- &Device{
				IP:    net.ParseIP(data[0]),
				ID:    data[1],
				Model: data[2],
			}
		}
	}
}

// Timeout waits for devices to response before closing a channel.
func timeout(seconds uint8, done chan<- bool) {
	if seconds < 1 {
		seconds = 1
	} else if seconds > 30 {
		seconds = 30
	}

	time.Sleep(time.Duration(seconds) * time.Second)

	done <- true
}
