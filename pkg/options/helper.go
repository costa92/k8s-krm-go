package options

import (
	"fmt"
	"net"
	"strings"

	netutils "k8s.io/utils/net"
)

// Define unit constant.
const (
	_   = iota // ignore onex.iota
	KiB = 1 << (10 * iota)
	MiB
	GiB
	TiB
)

func join(prefixs ...string) string {
	joined := strings.Join(prefixs, ".")
	if joined != "" {
		joined += "."
	}
	return joined
}

func ValidateAddress(addr string) error {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return err
	}
	if host != "" && netutils.ParseIPSloppy(host) == nil {
		return fmt.Errorf("%q is not a valid IP address", host)
	}
	if _, err := netutils.ParsePort(port, true); err != nil {
		return fmt.Errorf("%q is not a valid number", port)
	}

	return nil
}

// CreateListener create net listener by given address and returns it and port.
func CreateListener(addr string) (net.Listener, int, error) {
	network := "tcp"

	ln, err := net.Listen(network, addr)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to listen on %v: %w", addr, err)
	}

	// get port
	tcpAddr, ok := ln.Addr().(*net.TCPAddr)
	if !ok {
		_ = ln.Close()

		return nil, 0, fmt.Errorf("invalid listen address: %q", ln.Addr().String())
	}

	return ln, tcpAddr.Port, nil
}
