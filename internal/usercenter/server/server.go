package server

import (
	"context"
	"encoding/json"
	"github.com/costa92/k8s-krm-go/internal/usercenter/locales"
	"github.com/costa92/k8s-krm-go/pkg/i18n"
	"github.com/costa92/k8s-krm-go/pkg/log"
	krtlog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"golang.org/x/text/language"

	i18nmw "github.com/costa92/k8s-krm-go/internal/pkg/middleware/i18n"
)

// ProviderSet defines a wire provider set.
var ProviderSet = wire.NewSet(NewServers, NewHTTPServer, NewGRPCServer, NewMiddlewares)

func NewServers(hs *http.Server, gs *grpc.Server) []transport.Server {
	return []transport.Server{hs, gs}
}

func NewMiddlewares(logger krtlog.Logger) []middleware.Middleware {
	return []middleware.Middleware{
		recovery.Recovery(
			recovery.WithHandler(func(ctx context.Context, rq, err any) error {
				data, _ := json.Marshal(rq)
				log.C(ctx).Errorw(err.(error), "Catching a panic", "rq", string(data))
				return nil
			}),
		),
		i18nmw.Translator(i18n.WithLanguage(language.English), i18n.WithFS(locales.Locales)),
		logging.Server(logger),
	}
}
