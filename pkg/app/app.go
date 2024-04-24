package app

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type App struct {
	name        string               // Name of the app
	shortDesc   string               // Short description of the app
	description string               // Description of the app
	run         RunFunc              // Function that runs the app
	cmd         *cobra.Command       // Cobra command
	args        cobra.PositionalArgs // Positional arguments

	healthCheckFunc HealthCheckFunc // Function that checks the health of the app
	// + options
	options CliOptions
	// + optional
	silence bool

	// + optional
	noConfig bool

	// watching and reloading config files
	// + optional
	watch bool
}

// RunFunc is a function that runs the app
type RunFunc func() error

// HealthCheckFunc is a function that checks the health of the app
type HealthCheckFunc func() error

type Option func(*App)

func WithOptions(options CliOptions) Option {
	return func(a *App) {
		a.options = options
	}
}

func WithRunFunc(run RunFunc) Option {
	return func(a *App) {
		a.run = run
	}
}

func WithDescription(description string) Option {
	return func(a *App) {
		a.description = description
	}
}

func WithHealthCheckFunc(healthCheckFunc HealthCheckFunc) Option {
	return func(a *App) {
		a.healthCheckFunc = healthCheckFunc
	}
}

func WithDefaultHealthCheckFunc() Option {
	fn := func() HealthCheckFunc {
		return func() error {
			return nil
		}
	}
	return WithHealthCheckFunc(fn())
}

func WithSilence() Option {
	return func(a *App) {
		a.silence = true
	}
}

func WithNoConfig() Option {
	return func(a *App) {
		a.noConfig = true
	}
}

func WithValidArgs(args cobra.PositionalArgs) Option {
	return func(a *App) {
		a.args = args
	}
}

func WithDefaultValidArgs() Option {
	return func(a *App) {
		a.args = func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		}
	}
}

func WithWatch() Option {
	return func(a *App) {
		a.watch = true
	}
}

func NewApp(name, shortDesc string, opts ...Option) *App {
	app := &App{
		name:      name,
		run:       func() error { return nil },
		shortDesc: shortDesc,
	}
	for _, opt := range opts {
		opt(app)
	}

	//  a.buildCommand()
	return app
}

func (a *App) buildCommand() {
	cmd := &cobra.Command{
		Use:   a.name,
		Short: a.shortDesc,
		Long:  a.description,
		RunE:  a.runCommand,
		Args:  a.args,
	}

	a.cmd = cmd
}

// Run runs the app
func (a *App) runCommand(cmd *cobra.Command, args []string) error {
	// todo print the version
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return err
	}

	if a.options != nil {
		if err := a.options.Validate(); err != nil {
			return err
		}
	}

	return nil
}
