package options

import "github.com/spf13/pflag"

type IOptions interface {
	// Validate validates all the required options. It can also used to complete options if needed.
	Validate() []error

	// AddFlags adds flags related to given flagset.
	AddFlags(fs *pflag.FlagSet, prefixs ...string)
}
