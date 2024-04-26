package bootstrap

import "os"

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
