package app

import (
	"github.com/costa92/k8s-krm-go/cmd/usercenter/app/options"
	"github.com/costa92/k8s-krm-go/internal/usercenter"
	"github.com/costa92/k8s-krm-go/pkg/app"
)

// Define the description of the command.
const commandDesc = `The usercenter server is used to manage users, keys, fees, etc.`

// NewApp App is the main application
func NewApp(name string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp(
		name,
		"Launch the usercenter service",
		app.WithDescription("Launch the usercenter service"),
		app.WithOptions(opts),
		app.WithDefaultValidArgs(),
		app.WithRunFunc(run(opts)),
	)
	return application
}

func run(opts *options.Options) app.RunFunc {
	return func() error {
		cfg, err := opts.Config()
		if err != nil {
			return err
		}
		return Run(cfg, nil)
	}
}

// Run the application
func Run(cfg *usercenter.Config, stopCh <-chan struct{}) error {
	server, err := cfg.Complete().New(stopCh)
	if err != nil {
		return err
	}
	return server.Run(stopCh)
}
