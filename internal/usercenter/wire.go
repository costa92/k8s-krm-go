// Copyright All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
//

//go:build wireinject
// +build wireinject

package usercenter

//go:generate go run github.com/google/wire/cmd/wire

import (
	"github.com/costa92/k8s-krm-go/internal/pkg/bootstrap"
	"github.com/costa92/k8s-krm-go/internal/usercenter/server"
	"github.com/costa92/k8s-krm-go/internal/usercenter/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

// wireApp creates a new kratos application with the necessary dependencies.
func wireApp(
	bootstrap.AppInfo,
	*server.Config,
) (*kratos.App, func(), error) {
	wire.Build(
		bootstrap.ProviderSet,
		server.ProviderSet,
		service.ProviderSet,
	)
	return nil, nil, nil
}
