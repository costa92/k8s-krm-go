package bootstrap

import (
	"os"

	"github.com/go-kratos/kratos/v2"
	krtlog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/google/wire"
)

// ProviderSet is the provider set for the bootstrap
var ProviderSet = wire.NewSet(wire.Struct(new(AppConfig), "*"), NewLogger, NewApp)

// AppInfo is the information of the application
type AppInfo struct {
	ID       string
	Name     string
	Version  string
	Metadata map[string]string
}

func NewAppInfo(id, name, version string) AppInfo {
	if id == "" {
		id, _ = os.Hostname()
	}
	// Return the AppInfo struct
	return AppInfo{
		ID:       id,
		Name:     name,
		Version:  version,
		Metadata: map[string]string{},
	}
}

// AppConfig Config is the configuration of the application
type AppConfig struct {
	Info   AppInfo
	Logger krtlog.Logger
	//Registrar registry.Registrar
}

// NewApp creates a new kratos application
func NewApp(c AppConfig, servers ...transport.Server) *kratos.App {
	return kratos.New(
		kratos.ID(c.Info.ID+"."+c.Info.Name),
		kratos.Name(c.Info.Name),
		kratos.Version(c.Info.Version),
		kratos.Metadata(c.Info.Metadata),
		kratos.Logger(c.Logger),
		//kratos.Registrar(c.Registrar),
		kratos.Server(servers...),
	)
}
