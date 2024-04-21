package app

import "github.com/spf13/cobra"

type App struct {
	name        string               // Name of the app
	shortDesc   string               // Short description of the app
	description string               // Description of the app
	run         RunFunc              // Function that runs the app
	cmd         *cobra.Command       // Cobra command
	args        cobra.PositionalArgs // Positional arguments

	healthCheckFunc HealthCheckFunc // Function that checks the health of the app

	options CliOptions
}

// RunFunc is a function that runs the app
type RunFunc func() error

// HealthCheckFunc is a function that checks the health of the app
type HealthCheckFunc func() error

type Option func(*App)
