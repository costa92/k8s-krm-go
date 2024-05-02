package usercenter

import (
	"github.com/costa92/k8s-krm-go/internal/pkg/bootstrap"
	"github.com/costa92/k8s-krm-go/internal/usercenter/server"
	"github.com/costa92/k8s-krm-go/pkg/db"
	"github.com/costa92/k8s-krm-go/pkg/log"
	genericoptions "github.com/costa92/k8s-krm-go/pkg/options"
	"github.com/go-kratos/kratos/v2"
	"github.com/jinzhu/copier"
	"os"
)

var (
	Name  = "usercenter"
	ID, _ = os.Hostname()
)

type Config struct {
	HTTPOptions  *genericoptions.HTTPOptions
	GRPCOption   *genericoptions.GRPCOptions
	MySQLOptions *genericoptions.MySQLOptions
}

// completedConfig is a Config that has been completed with the necessary information
type completedConfig struct {
	*Config
}

// Complete NewConfig returns a new Config for the usercenter
func (cfg *Config) Complete() completedConfig {
	return completedConfig{cfg}
}

func (c completedConfig) New(stopCh <-chan struct{}) (*Server, error) {
	appInfo := bootstrap.NewAppInfo(ID, Name, "v0.0.1")
	conf := &server.Config{
		HTTP: *c.HTTPOptions,
		GRPC: *c.GRPCOption,
	}

	var dbOptions db.MySQLOptions
	_ = copier.Copy(&dbOptions, c.MySQLOptions)

	log.Infow("usercenter config", "http", conf.HTTP, "grpc", conf.GRPC)
	app, cleanup, err := wireApp(appInfo, conf, &dbOptions)
	if err != nil {
		return nil, err
	}
	defer cleanup()
	return &Server{app: app}, nil
}

// Server is the usercenter server
type Server struct {
	app *kratos.App
}

// Run runs the usercenter server
func (s *Server) Run(stopCh <-chan struct{}) error {
	go func() {
		if err := s.app.Run(); err != nil {
			log.Errorf("Failed to run server: %v", err)
		}
	}()

	<-stopCh
	log.Infof("Gracefully shutting down server ...")

	if err := s.app.Stop(); err != nil {
		log.Errorf("Failed to stop server: %v", err)
	}
	return nil
}
