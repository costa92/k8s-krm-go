//go:build wireinject
// +build wireinject

package usercenter

//go:generate go run github.com/google/wire/cmd/wire

import (
	"github.com/costa92/k8s-krm-go/internal/pkg/bootstrap"
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

// wireApp creates a new kratos application with the necessary dependencies.
func wireApp(
	bootstrap.AppInfo,
) (*kratos.App, func(), error) {

	wire.Build(
		bootstrap.ProviderSet,
	)

	return nil, nil, nil
}
