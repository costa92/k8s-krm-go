package options

import (
	"github.com/costa92/k8s-krm-go/internal/usercenter"
	"github.com/costa92/k8s-krm-go/pkg/app"
	"github.com/costa92/k8s-krm-go/pkg/log"
	genericoptions "github.com/costa92/k8s-krm-go/pkg/options"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apiserver/pkg/util/feature"
	cliflag "k8s.io/component-base/cli/flag"
)

var _ app.CliOptions = (*Options)(nil)

type Options struct {
	// gRPC options for configuring gRPC related options.
	GRPCOptions *genericoptions.GRPCOptions `json:"grpc" mapstructure:"grpc"`
	// HTTP options for configuring HTTP related options.
	HTTPOptions *genericoptions.HTTPOptions `json:"http" mapstructure:"http"`
	// TLS options for configuring TLS related options.
	TLSOptions *genericoptions.TLSOptions `json:"tls" mapstructure:"tls"`
	// MySQL
	MySQLOptions *genericoptions.MySQLOptions `json:"mysql" mapstructure:"mysql"`
	// Redis
	RedisOptions *genericoptions.RedisOptions `json:"redis" mapstructure:"redis"`
	// Kafka options for configuring Kafka related options.
	KafkaOptions *genericoptions.KafkaOptions `json:"kafka" mapstructure:"kafka"`
	// jwt options for configuring jwt related options.
	JWTOptions *genericoptions.JWTOptions `json:"jwt" mapstructure:"jwt"`

	// Log options for configuring log related options.
	Log *log.Options `json:"log" mapstructure:"log"`
}

func NewOptions() *Options {
	return &Options{
		GRPCOptions:  genericoptions.NewGRPCOptions(),
		HTTPOptions:  genericoptions.NewHTTPOptions(),
		MySQLOptions: genericoptions.NewMySQLOption(),
		RedisOptions: genericoptions.NewRedisOptions(),
		TLSOptions:   genericoptions.NewTLSOptions(),
		KafkaOptions: genericoptions.NewKafkaOptions(),
		JWTOptions:   genericoptions.NewJWTOptions(),
		Log:          log.NewOptions(),
	}
}

func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.GRPCOptions.AddFlags(fss.FlagSet("grpc"))
	o.HTTPOptions.AddFlags(fss.FlagSet("http"))
	o.TLSOptions.AddFlags(fss.FlagSet("tls"))
	o.MySQLOptions.AddFlags(fss.FlagSet("mysql"))
	o.RedisOptions.AddFlags(fss.FlagSet("redis"))
	o.KafkaOptions.AddFlags(fss.FlagSet("kafka"))
	o.JWTOptions.AddFlags(fss.FlagSet("jwt"))
	o.Log.AddFlags(fss.FlagSet("log"))
	fs := fss.FlagSet("misc")
	//client.AddFlags(fs)
	feature.DefaultMutableFeatureGate.AddFlag(fs)
	return fss
}

func (o *Options) Complete() error {
	return nil
}

func (o *Options) Validate() error {
	var errs []error
	errs = append(errs, o.GRPCOptions.Validate()...)
	errs = append(errs, o.HTTPOptions.Validate()...)
	errs = append(errs, o.TLSOptions.Validate()...)
	errs = append(errs, o.MySQLOptions.Validate()...)
	errs = append(errs, o.RedisOptions.Validate()...)
	errs = append(errs, o.KafkaOptions.Validate()...)
	errs = append(errs, o.JWTOptions.Validate()...)
	errs = append(errs, o.Log.Validate()...)
	return utilerrors.NewAggregate(errs)
}

// ApplyTo applies the options to the usercenter config
func (o *Options) ApplyTo(c *usercenter.Config) error {
	c.HTTPOptions = o.HTTPOptions
	c.GRPCOption = o.GRPCOptions
	c.TLSOptions = o.TLSOptions
	c.MySQLOptions = o.MySQLOptions
	c.RedisOptions = o.RedisOptions
	c.KafkaOptions = o.KafkaOptions
	c.JWTOptions = o.JWTOptions
	return nil
}

// Config returns a new usercenter config
func (o *Options) Config() (*usercenter.Config, error) {
	c := &usercenter.Config{}
	if err := o.ApplyTo(c); err != nil {
		return nil, err
	}
	return c, nil
}
