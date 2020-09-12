package magichome

import (
	"fmt"
	"net"
)

// A Command represents a powered state and a device color.
type Command struct {
	on    []byte
	off   []byte
	color []byte
}

// Disover returns a list of Device's found on the local area network.
// Device's are returned using the order in which tthey are encoutnered.
func Discover(options DiscoverOptions) (*[]Device, error) {

	const port = "48899"
	const protcol = "udp4"

	if options.BroadcastAddr == "" {
		options.BroadcastAddr = DEFAULT_BROADCAST_ADDR
	}

	broadcastAddr, err := net.ResolveUDPAddr(protcol, fmt.Sprintf("%s:%s", options.BroadcastAddr, port))
	if err != nil {
		return nil, err
	}

	localAddr, err := net.ResolveUDPAddr(protcol, ":0")
	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUDP(protcol, localAddr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	received := make(chan *Device)
	done := make(chan bool, 1)
	var devices []Device

	go receive(conn, received)
	go send(conn, broadcastAddr)
	go timeout(options.Timeout, done)

	for {
		select {
		case device, ok := <-received:
			if ok {
				devices = append(devices, *device)
			}
		case <-done:
			done = nil
		}

		if done == nil {
			break
		}
	}

	close(received)

	return &devices, nil
}
