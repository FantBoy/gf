// Copyright 2017-2018 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gudp

import (
	"net"
)

// NewNetConn creates and returns a *net.UDPConn with given addresses.
func NewNetConn(remoteAddress string, localAddress ...string) (*net.UDPConn, error) {
	var err error
	var remoteAddr, localAddr *net.UDPAddr
	remoteAddr, err = net.ResolveUDPAddr("udp", remoteAddress)
	if err != nil {
		return nil, err
	}
	if len(localAddress) > 0 {
		localAddr, err = net.ResolveUDPAddr("udp", localAddress[0])
		if err != nil {
			return nil, err
		}
	}
	conn, err := net.DialUDP("udp", localAddr, remoteAddr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// Send writes data to <addr> using UDP connection and then closes the connection.
// Note that it is used for short connection usage.
func Send(addr string, data []byte, retry ...Retry) error {
	conn, err := NewConn(addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	return conn.Send(data, retry...)
}

// SendRecv writes data to <addr> using UDP connection, reads response and then closes the connection.
// Note that it is used for short connection usage.
func SendRecv(addr string, data []byte, receive int, retry ...Retry) ([]byte, error) {
	conn, err := NewConn(addr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return conn.SendRecv(data, receive, retry...)
}

// isTimeout checks whether given <err> is a timeout error.
func isTimeout(err error) bool {
	if err == nil {
		return false
	}
	if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
		return true
	}
	return false
}
