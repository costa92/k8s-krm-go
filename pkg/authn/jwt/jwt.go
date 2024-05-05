package jwt

import (
	"context"
	"github.com/costa92/k8s-krm-go/pkg/authn"
	"github.com/costa92/k8s-krm-go/pkg/i18n"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/golang-jwt/jwt/v4"
	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
	"time"
)

const (
	// Unauthorized is a constant that represents an unauthorized error.
	reason string = "Unauthorized"

	defaultKey string = "krm"
)

var (
	ErrTokenInvalid           = errors.Unauthorized(reason, "Token is invalid")
	ErrTokenExpired           = errors.Unauthorized(reason, "Token has expired")
	ErrTokenParseFail         = errors.Unauthorized(reason, "Fail to parse token")
	ErrUnSupportSigningMethod = errors.Unauthorized(reason, "Wrong signing method")
	ErrSignTokenFailed        = errors.Unauthorized(reason, "Fail to sign token")
)

// Define i18n messages.
var (
	MessageTokenInvalid           = &goi18n.Message{ID: "jwt.token.invalid", Other: ErrTokenInvalid.Error()}
	MessageTokenExpired           = &goi18n.Message{ID: "jwt.token.expired", Other: ErrTokenExpired.Error()}
	MessageTokenParseFail         = &goi18n.Message{ID: "jwt.token.parse.failed", Other: ErrTokenParseFail.Error()}
	MessageUnSupportSigningMethod = &goi18n.Message{ID: "jwt.wrong.signing.method", Other: ErrUnSupportSigningMethod.Error()}
	MessageSignTokenFailed        = &goi18n.Message{ID: "jwt.token.sign.failed", Other: ErrSignTokenFailed.Error()}
)

var defaultOptions = options{
	tokenType:     "Bearer",
	expired:       2 * time.Hour,
	signingMethod: jwt.SigningMethodHS256,
	signingKey:    []byte(defaultKey),
	keyfunc: func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrTokenInvalid
		}
		return []byte(defaultKey), nil
	},
}

type options struct {
	signingMethod jwt.SigningMethod // Signing method.
	signingKey    any               // Signing key.
	keyfunc       jwt.Keyfunc       // Key function.
	issuer        string            // Issuer.
	expired       time.Duration     // Expiration time.
	tokenType     string            // Token type.
	tokenHeader   map[string]any    // Token header.
}

type Option func(*options)

// WithSigningMethod sets the signing method.
func WithSigningMethod(method jwt.SigningMethod) Option {
	return func(o *options) {
		o.signingMethod = method
	}
}

// WithSigningKey sets the signing key.
func WithSigningKey(key any) Option {
	return func(o *options) {
		o.signingKey = key
	}
}

// WithKeyFunc sets the key function.
func WithKeyFunc(keyfunc jwt.Keyfunc) Option {
	return func(o *options) {
		o.keyfunc = keyfunc
	}
}

// WithIssuer sets the token issuer.
func WithIssuer(issuer string) Option {
	return func(o *options) {
		o.issuer = issuer
	}
}

// WithExpired sets the expiration time.
func WithExpired(expired time.Duration) Option {
	return func(o *options) {
		o.expired = expired
	}
}

// WithTokenType sets the token type.
func WithTokenType(tokenType string) Option {
	return func(o *options) {
		o.tokenType = tokenType
	}
}

// WithTokenHeader sets the token header.
func WithTokenHeader(header map[string]any) Option {
	return func(o *options) {
		o.tokenHeader = header
	}
}

func New(store Storer, opts ...Option) *JWTAuth {
	opt := defaultOptions
	for _, o := range opts {
		o(&opt)
	}
	return &JWTAuth{
		opts:  &opt,
		store: store,
	}
}

// JWTAuth is a jwt authentication.
type JWTAuth struct {
	opts  *options
	store Storer
}

func (a *JWTAuth) Issue(ctx context.Context, userID string) (authn.IToken, error) {
	now := time.Now()
	expiresAt := now.Add(a.opts.expired)

	token := jwt.NewWithClaims(a.opts.signingMethod, &jwt.RegisteredClaims{
		// Issuer = iss,令牌颁发者。它表示该令牌是由谁创建的
		Issuer: a.opts.issuer,
		// IssuedAt = iat,令牌颁发时的时间戳。它表示令牌是何时被创建的
		IssuedAt: jwt.NewNumericDate(now),
		// ExpiresAt = exp,令牌过期时间。它表示令牌何时会过期
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		// NotBefore = nbf,令牌的生效时间。它表示令牌在何时开始生效
		NotBefore: jwt.NewNumericDate(now),
		// Subject = sub,令牌的主题。它表示令牌的主题
		Subject: userID,
	})

	if a.opts.tokenHeader != nil {
		for k, v := range a.opts.tokenHeader {
			token.Header[k] = v
		}
	}
	refreshToken, err := token.SignedString(a.opts.signingKey)
	if err != nil {
		return nil, errors.Unauthorized(reason, err.Error())
	}
	tokenInfo := &tokenInfo{
		ExpiresAt: expiresAt.Unix(),
		Type:      a.opts.tokenType,
		Token:     refreshToken,
	}
	return tokenInfo, nil
}

// parseToken is used to parse the input refreshToken.
func (a *JWTAuth) parseToken(ctx context.Context, refreshToken string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &jwt.RegisteredClaims{}, a.opts.keyfunc)
	if err != nil {
		ve, ok := err.(*jwt.ValidationError)
		if !ok {
			return nil, errors.Unauthorized(reason, err.Error())
		}
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, errors.Unauthorized(reason, i18n.FromContext(ctx).LocalizeT(MessageTokenInvalid))
		}
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, errors.Unauthorized(reason, i18n.FromContext(ctx).LocalizeT(MessageTokenExpired))
		}
		return nil, errors.Unauthorized(reason, i18n.FromContext(ctx).LocalizeT(MessageTokenParseFail))
	}

	if !token.Valid {
		return nil, errors.Unauthorized(reason, i18n.FromContext(ctx).LocalizeT(MessageTokenInvalid))
	}

	if token.Method != a.opts.signingMethod {
		return nil, errors.Unauthorized(reason, i18n.FromContext(ctx).LocalizeT(MessageUnSupportSigningMethod))
	}
	return token.Claims.(*jwt.RegisteredClaims), nil
}

// ParseClaims is a method that parses the token and returns the claims.
func (a *JWTAuth) callStore(fn func(Storer) error) error {
	if store := a.store; store != nil {
		return fn(store)
	}
	return nil
}

// Destroy is used to destroy a token.
func (a *JWTAuth) Destroy(ctx context.Context, refreshToken string) error {
	claims, err := a.parseToken(ctx, refreshToken)
	if err != nil {
		return err
	}

	// If storage is set, put the unexpired token in
	store := func(store Storer) error {
		expired := time.Until(claims.ExpiresAt.Time)
		return store.Set(ctx, refreshToken, expired)
	}
	return a.callStore(store)
}

// ParseClaims parse the token and return the claims.
func (a *JWTAuth) ParseClaims(ctx context.Context, refreshToken string) (*jwt.RegisteredClaims, error) {
	if refreshToken == "" {
		return nil, errors.Unauthorized(reason, i18n.FromContext(ctx).LocalizeT(MessageTokenInvalid))
	}

	claims, err := a.parseToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	store := func(store Storer) error {
		exists, err := store.Check(ctx, refreshToken)
		if err != nil {
			return err
		}

		if exists {
			return errors.Unauthorized(reason, i18n.FromContext(ctx).LocalizeT(MessageTokenInvalid))
		}

		return nil
	}

	if err := a.callStore(store); err != nil {
		return nil, err
	}

	return claims, nil
}

// Release used to release the requested resources.
func (a *JWTAuth) Release() error {
	return a.callStore(func(store Storer) error {
		return store.Close()
	})
}
