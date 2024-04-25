package options

import (
	"github.com/costa92/k8s-krm-go/pkg/app"
	cliflag "k8s.io/component-base/cli/flag"
)

var _ app.CliOptions = (*Options)(nil)

type Options struct {
}

func NewOptions() *Options {
	return &Options{}
}

func (o Options) Flags() cliflag.NamedFlagSets {
	//TODO implement me
	panic("implement me")
}

func (o Options) Complete() error {
	//TODO implement me
	panic("implement me")
}

func (o Options) Validate() error {
	//TODO implement me
	panic("implement me")
}

func (o *Options) Config() (interface{}, error) {
	//TODO implement me
	panic("implement me")
}
