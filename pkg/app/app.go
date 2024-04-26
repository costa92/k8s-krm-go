package app

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/costa92/k8s-krm-go/pkg/log"
	"k8s.io/component-base/cli"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/term"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type App struct {
	name        string               // Name of the app
	shortDesc   string               // Short description of the app
	description string               // Description of the app
	run         RunFunc              // Function that runs the app
	cmd         *cobra.Command       // Cobra command
	args        cobra.PositionalArgs // Positional arguments

	healthCheckFunc HealthCheckFunc // Function that checks the health of the app
	// + options
	options CliOptions
	// + optional
	silence bool

	// + optional
	noConfig bool

	// watching and reloading config files
	// + optional
	watch bool
}

// RunFunc is a function that runs the app
type RunFunc func() error

// HealthCheckFunc is a function that checks the health of the app
type HealthCheckFunc func() error

type Option func(*App)

func WithOptions(options CliOptions) Option {
	return func(a *App) {
		a.options = options
	}
}

func WithRunFunc(run RunFunc) Option {
	return func(a *App) {
		a.run = run
	}
}

func WithDescription(description string) Option {
	return func(a *App) {
		a.description = description
	}
}

func WithHealthCheckFunc(healthCheckFunc HealthCheckFunc) Option {
	return func(a *App) {
		a.healthCheckFunc = healthCheckFunc
	}
}

func WithDefaultHealthCheckFunc() Option {
	fn := func() HealthCheckFunc {
		return func() error {
			return nil
		}
	}
	return WithHealthCheckFunc(fn())
}

func WithSilence() Option {
	return func(a *App) {
		a.silence = true
	}
}

func WithNoConfig() Option {
	return func(a *App) {
		a.noConfig = true
	}
}

func WithValidArgs(args cobra.PositionalArgs) Option {
	return func(a *App) {
		a.args = args
	}
}

func WithDefaultValidArgs() Option {
	return func(a *App) {
		a.args = func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		}
	}
}

func WithWatch() Option {
	return func(a *App) {
		a.watch = true
	}
}

func NewApp(name, shortDesc string, opts ...Option) *App {
	a := &App{
		name:      name,
		run:       func() error { return nil },
		shortDesc: shortDesc,
	}
	for _, opt := range opts {
		opt(a)
	}

	a.buildCommand()

	return a
}

func (a *App) buildCommand() {
	cmd := &cobra.Command{
		Use:   formatBaseName(a.name),
		Short: a.shortDesc,
		Long:  a.description,
		RunE:  a.runCommand,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		Args: a.args,
	}

	if !cmd.SilenceUsage { // 如果不是静默模式
		cmd.SilenceUsage = true
		// SetFlagErrorFunc sets the function that will be called if an error occurs while parsing the flags.
		cmd.SetFlagErrorFunc(func(c *cobra.Command, err error) error { // 设置 flag 解析错误时的回调函数
			c.SilenceUsage = false
			return err
		})
	}

	cmd.SilenceErrors = true
	cmd.SetOutput(os.Stdout)
	cmd.SetErr(os.Stderr)        // SetErr sets the destination for error messages written by the command.
	cmd.Flags().SortFlags = true // SortFlags sets the flag sorting function.

	var fss cliflag.NamedFlagSets
	if a.options != nil {
		fss = a.options.Flags()
	}
	// todo version
	if !a.noConfig { // 如果不是 noConfig 模式
		AddConfigFlag(fss.FlagSet("global"), a.name, a.watch)
	}
	// Add the flags to the command
	for _, f := range fss.FlagSets {
		cmd.Flags().AddFlagSet(f) // AddFlagSet adds all the flags in the FlagSet to the command.
	}
	// SetUsageFunc sets the function that will be called if an error occurs while parsing the flags.
	cols, _, _ := term.TerminalSize(cmd.OutOrStdout()) // TerminalSize returns the dimensions of the given terminal.
	cliflag.SetUsageAndHelpFunc(cmd, fss, cols)        //

	a.cmd = cmd
}

func (a *App) Run() {
	os.Exit(cli.Run(a.cmd))
}

// Run runs the app
func (a *App) runCommand(cmd *cobra.Command, args []string) error {
	// todo print the version
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return err
	}
	if a.options != nil {
		// complete the configuration
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			return err
		}
		// complete the configuration
		if err := a.options.Complete(); err != nil {
			return err
		}
		// validate the configuration
		if err := a.options.Validate(); err != nil {
			return err
		}
	}
	// 初始化日志
	log.Init(logOptions())
	defer log.Sync() // sync 将缓存中的日志刷新到硬盘文件

	if !a.silence {
		log.Infow("Starting the app", "name", a.name)
		log.Infow("Golang settings", "GOGC", os.Getenv("GOGC"), "GOMAXPROCS", os.Getenv("GOMAXPROCS"), "GOTRACEBACK", os.Getenv("GOTRACEBACK"))
		if !a.noConfig {
			PrintConfig() // 打印配置
		} else if a.options != nil {
			cliflag.PrintFlags(cmd.Flags()) // 打印 flags
		}
	}

	if a.healthCheckFunc != nil {
		if err := a.healthCheckFunc(); err != nil {
			return err
		}
	}
	return a.run()
}

func (a *App) Command() *cobra.Command {
	return a.cmd
}

func formatBaseName(name string) string {
	if runtime.GOOS == "windows" {
		name = strings.ToLower(name)
		name = strings.TrimSuffix(name, ".exe")
	}
	return name
}

// logOptions 从 viper 中读取日志配置，构建 `*log.Options` 并返回.
// 注意：`viper.Get<Type>()` 中 key 的名字需要使用 `.` 分割，以跟 YAML 中保持相同的缩进.
func logOptions() *log.Options {
	return &log.Options{
		DisableCaller:     viper.GetBool("log.disable-caller"),
		Level:             viper.GetString("log.level"),
		DisableStacktrace: viper.GetBool("log.disable-stacktrace"),
		Format:            viper.GetString("log.format"),
		OutputPaths:       viper.GetStringSlice("log.output-paths"),
		EnableColor:       viper.GetBool("log.enable-color"),
	}
}

// 从 viper 中读取配置，构建 `*log.Options` 并返回.
func init() {
	viper.SetDefault("log.disable-caller", false)
	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.disable-stacktrace", false)
	viper.SetDefault("log.format", "json")
	viper.SetDefault("log.output-paths", []string{"stdout"})
	viper.SetDefault("log.enable-color", false)
}
