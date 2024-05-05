package server

import (
	"context"
	"github.com/costa92/k8s-krm-go/internal/usercenter/service"
	"github.com/costa92/k8s-krm-go/pkg/api/usercenter/v1"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/kratos/v2/transport/http/pprof"
	"github.com/go-kratos/swagger-api/openapiv2"
	"github.com/gorilla/handlers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewWhiteListMatcher() selector.MatchFunc {
	whitelist := make(map[string]struct{})
	whitelist[v1.OperationUserServiceLogin] = struct{}{}
	whitelist[v1.OperationUserServiceLogout] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whitelist[operation]; ok {
			return false
		}
		return true
	}
}

func NewHTTPServer(c *Config, gw *service.UserCenterService, middlewares []middleware.Middleware) *http.Server {
	opts := []http.ServerOption{
		http.Middleware(middlewares...),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{
				"X-Requested-With",
				"Content-Type",
				"Authorization",
				"X-Idempotent-ID",
			}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowCredentials()),
		),
	}

	if c.HTTP.Network != "" {
		opts = append(opts, http.Network(c.HTTP.Network))
	}
	if c.HTTP.Timeout != 0 {
		opts = append(opts, http.Timeout(c.HTTP.Timeout))
	}
	if c.HTTP.Addr != "" {
		opts = append(opts, http.Address(c.HTTP.Addr))
	}

	srv := http.NewServer(opts...)
	h := openapiv2.NewHandler()
	srv.HandlePrefix("/openapi/", h)
	srv.Handle("/metrics", promhttp.Handler())
	srv.Handle("", pprof.NewHandler())

	v1.RegisterUserServiceHTTPServer(srv, gw)
	return srv
}
