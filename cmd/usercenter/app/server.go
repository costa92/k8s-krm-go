package app

import (
	"github.com/costa92/k8s-krm-go/cmd/usercenter/app/options"
	"github.com/costa92/k8s-krm-go/internal/usercenter"
	"github.com/costa92/k8s-krm-go/pkg/app"
	genericapiserver "k8s.io/apiserver/pkg/server"
)

// Define the description of the command.
const commandDesc = `The usercenter server is used to manage users, keys, fees, etc.`

// NewApp App is the main application
func NewApp(name string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp(name, "Launch the usercenter service",
		app.WithDescription(commandDesc),
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
		return Run(cfg, genericapiserver.SetupSignalHandler())
	}
}

// Run the application
func Run(cfg *usercenter.Config, stopCh <-chan struct{}) error {
	if cfg != nil {
		server, err := cfg.Complete().New(stopCh)
		if err != nil {
			return err
		}
		return server.Run(stopCh)
	}
	return nil
}
