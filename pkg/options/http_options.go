package options

import (
	"github.com/spf13/pflag"
	"time"
)

var _ IOptions = (*HTTPOptions)(nil)

type HTTPOptions struct {
	Network string        `json:"network" mapstructure:"network"`
	Addr    string        `json:"addr" mapstructure:"addr"`
	Timeout time.Duration `json:"timeout" mapstructure:"timeout"`
}

func NewHTTPOptions() *HTTPOptions {
	return &HTTPOptions{
		Network: "tcp",
		Addr:    ":8083",
		Timeout: 30 * time.Second,
	}
}

func (o *HTTPOptions) Validate() []error {
	if o == nil {
		return nil
	}
	var errors []error
	if err := ValidateAddress(o.Addr); err != nil {
		errors = append(errors, err)
	}
	return errors
}

func (o *HTTPOptions) AddFlags(fs *pflag.FlagSet, prefix ...string) {
	// Apply the options to the http server
	fs.StringVar(&o.Network, "http.network", o.Network, "The network type for the http server")
	fs.StringVar(&o.Addr, "http.addr", o.Addr, "The address for the http server")
	fs.DurationVar(&o.Timeout, "http.timeout", o.Timeout, "The timeout for the http server")
}

// Complete fills in any fields not set that are required to have valid data.
func (o *HTTPOptions) Complete() error {
	return nil
}
