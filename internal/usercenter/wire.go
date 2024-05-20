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
	"github.com/costa92/k8s-krm-go/internal/pkg/validation"
	"github.com/costa92/k8s-krm-go/internal/usercenter/auth"
	"github.com/costa92/k8s-krm-go/internal/usercenter/biz"
	"github.com/costa92/k8s-krm-go/internal/usercenter/server"
	"github.com/costa92/k8s-krm-go/internal/usercenter/service"
	"github.com/costa92/k8s-krm-go/internal/usercenter/store"
	usercentervalidation "github.com/costa92/k8s-krm-go/internal/usercenter/validation"
	"github.com/costa92/k8s-krm-go/pkg/db"
	genericoptions "github.com/costa92/k8s-krm-go/pkg/options"
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

// wireApp creates a new kratos application with the necessary dependencies.
func wireApp(
	bootstrap.AppInfo,
	*server.Config,
	*db.MySQLOptions,
	*genericoptions.JWTOptions,
	*genericoptions.RedisOptions,
	*genericoptions.KafkaOptions,
) (*kratos.App, func(), error) {
	wire.Build(
		bootstrap.ProviderSet,
		server.ProviderSet,
		store.ProviderSet,
		db.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		auth.ProviderSet,
		store.SetterProviderSet,
		NewAuthenticator,
		validation.ProviderSet,
		usercentervalidation.ProviderSet,
	)
	return nil, nil, nil
}
