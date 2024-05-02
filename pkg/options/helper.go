package options

import (
	"fmt"
	"net"
	"strings"

	netutils "k8s.io/utils/net"
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
