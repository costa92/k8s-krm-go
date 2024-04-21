package log

import (
	"go.uber.org/zap/zapcore"

	"github.com/spf13/pflag"
)

type Options struct {
	// DisableCaller is the caller level. Default is false.
	DisableCaller bool `json:"disable-caller,omitempty" mapstructure:"disable-caller"`
	// DisableStacktrace is the stacktrace level. Default is none.
	DisableStacktrace bool `json:"disable-stacktrace,omitempty" mapstructure:"disable-stacktrace"`

	EnableColor bool   `json:"enable-color" mapstructure:"enable-color"`
	Level       string `json:"level,omitempty" mapstructure:"level"`

	Format      string   `json:"format,omitempty" mapstructure:"format"`
	OutputPaths []string `json:"output-paths,omitempty" mapstructure:"output-paths"`
}

func NewOptions() *Options {
	return &Options{
		Level:       zapcore.InfoLevel.String(),
		Format:      "console",
		OutputPaths: []string{"stdout"},
	}
}

// Validate validates the Options
func (o *Options) Validate() []error {
	var errs []error
	return errs
}

func (o *Options) AddFlags(fs *pflag.FlagSet) *Options {
	fs.BoolVar(&o.DisableCaller, "log.disable-caller", o.DisableCaller, "disable caller")
	fs.BoolVar(&o.DisableStacktrace, "log.disable-stacktrace", o.DisableStacktrace, "disable stacktrace")
	fs.BoolVar(&o.EnableColor, "log.enable-color", o.EnableColor, "enable color")
	fs.StringVar(&o.Level, "log.level", o.Level, "log level")
	fs.StringVar(&o.Format, "log.format", o.Format, "log format")
	fs.StringSliceVar(&o.OutputPaths, "log.output-paths", o.OutputPaths, "log output paths")

	return o
}
