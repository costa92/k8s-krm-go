package app

import (
	"github.com/costa92/k8s-krm-go/cmd/usercenter/app/options"
	"github.com/costa92/k8s-krm-go/pkg/app"
)

// NewApp App is the main application
func NewApp(name string) *app.App {
	_ = options.NewOptions()
	return &app.App{
		//options: opts,
	}
}
