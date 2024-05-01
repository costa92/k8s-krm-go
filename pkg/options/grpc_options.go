package options

import (
	"github.com/spf13/pflag"
	"time"
)

var _ IOptions = (*GRPCOptions)(nil)

type GRPCOptions struct {
	// Network with server network.
	Network string `json:"network" mapstructure:"network"`
	// Address with server address.
	Addr string `json:"addr" mapstructure:"addr"`
	// Timeout with server timeout. Used by grpc client side.
	Timeout time.Duration `json:"timeout" mapstructure:"timeout"`
}

func NewGRPCOptions() *GRPCOptions {
	return &GRPCOptions{
		Network: "tcp",
		Addr:    "",
		Timeout: 30 * time.Second,
	}
}

// Validate is used to parse and validate the parameters entered by the user at
func (o *GRPCOptions) Validate() []error {
	var errors []error
	if err := ValidateAddress(o.Addr); err != nil {
		errors = append(errors, err)
	}
	return errors
}

// AddFlags adds flags related to features for a specific api server to the
func (o *GRPCOptions) AddFlags(fs *pflag.FlagSet, prefixs ...string) {
	fs.StringVar(&o.Network, "grpc.network", o.Network, "Specify the network for the gRPC server.")
	fs.StringVar(&o.Addr, "grpc.addr", o.Addr, "Specify the gRPC server bind address and port.")
	fs.DurationVar(&o.Timeout, "grpc.timeout", o.Timeout, "Timeout for server connections.")
}
