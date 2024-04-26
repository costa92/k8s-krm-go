package usercenter

import (
	"os"

	"github.com/costa92/k8s-krm-go/pkg/log"
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
	return nil, nil
}

type Server struct {
	//TODO implement me
}

// Run runs the usercenter server
func (s *Server) Run(stopCh <-chan struct{}) error {
	log.Infof("Gracefully shutting down server ...")
	return nil
}
