package usercenter

import (
	"os"

	"github.com/costa92/k8s-krm-go/internal/pkg/bootstrap"
	"github.com/costa92/k8s-krm-go/pkg/log"
	"github.com/go-kratos/kratos/v2"
)

var (
	Name  = "usercenter"
	ID, _ = os.Hostname()
)

type Config struct {
	//TODO implement me
}

// completedConfig is a Config that has been completed with the necessary information
type completedConfig struct {
	*Config
}

// NewConfig returns a new Config for the usercenter
func (cfg *Config) Complete() completedConfig {
	return completedConfig{cfg}
}

func (c completedConfig) New(stopCh <-chan struct{}) (*Server, error) {
	appInfo := bootstrap.NewAppInfo(ID, Name, "v0.0.1")
	app, cleanup, err := wireApp(appInfo)
	if err != nil {
		return nil, err
	}
	defer cleanup()
	return &Server{app: app}, nil
}

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
