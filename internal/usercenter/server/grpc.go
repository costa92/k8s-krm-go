package server

import (
	"github.com/costa92/k8s-krm-go/internal/usercenter/service"
	v1 "github.com/costa92/k8s-krm-go/pkg/api/usercenter/v1"
	"github.com/costa92/k8s-krm-go/pkg/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewGRPCServer(c *Config, uc *service.UserCenterService, middlewares []middleware.Middleware) *grpc.Server {
	opts := []grpc.ServerOption{
		// grpc.WithDiscovery(nil),
		// grpc.WithEndpoint("discovery:///matrix.creation.service.grpc"),
		// Define the middleware chain with variable options.
		grpc.Middleware(middlewares...),
	}
	if c.GRPC.Network != "" {
		opts = append(opts, grpc.Network(c.GRPC.Network))
	}
	if c.GRPC.Timeout != 0 {
		opts = append(opts, grpc.Timeout(c.GRPC.Timeout))
	}
	if c.GRPC.Addr != "" {
		opts = append(opts, grpc.Address(c.GRPC.Addr))
	}
	log.Infow("grpc server", "addr", c.GRPC.Addr, "network", c.GRPC.Network, "timeout", c.GRPC.Timeout)
	srv := grpc.NewServer(opts...)
	v1.RegisterUserServiceServer(srv, uc)
	return srv
}
