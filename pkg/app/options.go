package app

import (
	"path/filepath"
	"strings"

	"k8s.io/client-go/util/homedir"
	cliflag "k8s.io/component-base/cli/flag"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Option is a function that configures an App
const configFlagName = "config"

type CliOptions interface {
	// Flags returns the flags for the app
	Flags() cliflag.NamedFlagSets
	// Complete completes the configuration of the app
	Complete() error

	// Validate validates the configuration of the app
	Validate() error
}

var cfgFile string

func AddConfigFlag(fs *pflag.FlagSet, name string, watch bool) {
	// Add the flag to the FlagSet
	fs.AddFlag(pflag.Lookup(configFlagName)) // HL

	// Bind the flag to viper
	viper.AutomaticEnv()

	// Set the env prefix
	viper.SetEnvPrefix(strings.ReplaceAll(strings.ToUpper(name), "-", "_"))
	// Set the env key replacer
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	// Bind the flag to viper
	cobra.OnInitialize(func() {
		if cfgFile != "" {
			viper.SetConfigFile(cfgFile)
		} else {
			viper.AddConfigPath(".")
			if names := strings.Split(name, "-"); len(names) > 1 {
				viper.AddConfigPath(filepath.Join(homedir.HomeDir(), ","+names[0]))
				viper.AddConfigPath(filepath.Join("/etc", "."+names[0]))
			}
			viper.SetConfigName(name)
		}

		if err := viper.ReadInConfig(); err != nil {
		}
	})
}
