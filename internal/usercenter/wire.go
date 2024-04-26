package usercenter

import (
	"github.com/costa92/k8s-krm-go/internal/pkg/bootstrap"
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

type runFunc func()

func wireApp(bootstrap.AppInfo) (*kratos.App, runFunc, error) {
	wire.Build(
		bootstrap.ProviderSet,
	)
	return nil, nil, nil
}
