package options

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/spf13/pflag"
	"time"
)

var _ IOptions = (*JWTOptions)(nil)

type JWTOptions struct {
	Key     string        `json:"key" mapstructure:"key"`
	Expired time.Duration `json:"expired" mapstructure:"expired"`
	// MaxRefresh is the maximum time that a client can refresh their token.
	MaxRefresh time.Duration `json:"max-refresh" mapstructure:"max-refresh"`
	// SigningMethod is the method used to sign the JWT token.
	SigningMethod string `json:"signing-method" mapstructure:"signing-method"`
}

func NewJWTOptions() *JWTOptions {
	return &JWTOptions{
		Key:           "krm(#)666",
		Expired:       2 * time.Hour,
		MaxRefresh:    2 * time.Hour,
		SigningMethod: "HS512",
	}
}

func (s *JWTOptions) Validate() []error {
	var errs []error
	if !govalidator.StringLength(s.Key, "6", "32") {
		errs = append(errs, fmt.Errorf("--jwt.key must larger than 5 and little than 33"))
	}
	return errs
}

// AddFlags adds flags related to features for a specific api server to the
// specified FlagSet.
func (s *JWTOptions) AddFlags(fs *pflag.FlagSet, prefixs ...string) {
	if fs == nil {
		return
	}

	// fs.StringVar(&s.Realm, "jwt.realm", s.Realm, "Realm name to display to the user.")
	fs.StringVar(&s.Key, "jwt.key", s.Key, "Private key used to sign jwt token.")
	fs.DurationVar(&s.Expired, "jwt.expired", s.Expired, "JWT token expiration time.")
	fs.DurationVar(&s.MaxRefresh, "jwt.max-refresh", s.MaxRefresh, ""+
		"This field allows clients to refresh their token until MaxRefresh has passed.")
	fs.StringVar(&s.SigningMethod, "jwt.signing-method", s.SigningMethod, "JWT token signature method.")
}
