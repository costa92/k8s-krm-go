package auth

import (
	"context"
	known "github.com/costa92/k8s-krm-go/internal/pkg/known/usercenter"
	"github.com/costa92/k8s-krm-go/internal/usercenter/model"
	"github.com/costa92/k8s-krm-go/pkg/authn"
	jwtauthn "github.com/costa92/k8s-krm-go/pkg/authn/jwt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/wire"
	"k8s.io/utils/lru"
	"time"
)

const (
	// reasonUnauthorized holds the error reason.
	reasonUnauthorized string = "Unauthorized"
)

// AuthnProviderSet is authn providers.
var AuthnProviderSet = wire.NewSet(NewAuthn, wire.Bind(new(AuthnInterface), new(*authnImpl)))

var (
	// ErrMissingKID is returned when the token format is invalid and the kid field is missing in the token header.
	ErrMissingKID = errors.Unauthorized(reasonUnauthorized, "Invalid token format: missing kid field in header")
	// ErrSecretDisabled is returned when the SecretID is disabled.
	ErrSecretDisabled = errors.Unauthorized(reasonUnauthorized, "SecretID is disabled")
)

type AuthnInterface interface {
	// Sign signs the user in and returns a token.
	Sign(ctx context.Context, userID string) (authn.IToken, error)
	// Verify verifies the token and returns the user ID.
	Verify(accessToken string) (string, error)
}

type TemporarySecretSetter interface {
	Get(ctx context.Context, secretID string) (*model.SecretM, error)
	Set(ctx context.Context, userID string, expires int64) (*model.SecretM, error)
}

// authnImpl is an implementation of the AuthnInterface.
type authnImpl struct {
	setter  TemporarySecretSetter
	secrets *lru.Cache // LRU cache for storing temporary secrets.
}

// Ensure authnImpl implements AuthnInterface.
var _ AuthnInterface = (*authnImpl)(nil)

// NewAuthn creates a new instance of the AuthnInterface.
func NewAuthn(setter TemporarySecretSetter) (*authnImpl, error) {
	l := lru.New(known.DefaultLRUSize)
	return &authnImpl{setter: setter, secrets: l}, nil
}

// Sign signs the user in and returns a token.
func (a *authnImpl) Sign(ctx context.Context, userID string) (authn.IToken, error) {
	expires := time.Now().Add(known.AccessTokenExpire).Unix()
	// Set the temporary secret key pair.
	secret, err := a.setter.Set(ctx, userID, expires)
	if err != nil {
		return nil, err
	}
	// Create the options for the token.
	opts := []jwtauthn.Option{
		jwtauthn.WithSigningMethod(jwt.SigningMethodHS512),
		jwtauthn.WithIssuer("sercenter"),
		jwtauthn.WithTokenHeader(map[string]interface{}{"kid": secret.SecretID}),
		jwtauthn.WithExpired(known.AccessTokenExpire),
		jwtauthn.WithSigningKey([]byte(secret.SecretKey)),
	}

	j, err := jwtauthn.New(nil, opts...).Sign(ctx, userID)
	if err != nil {
		return nil, err
	}
	return j, nil
}

// Verify verifies the token and returns the user ID.
func (a *authnImpl) Verify(accessToken string) (string, error) {
	var secret *model.SecretM
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		// Validate the alg is HMAC signature
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", jwtauthn.ErrUnSupportSigningMethod
		}

		kid, ok := token.Header["kid"].(string)
		if !ok {
			return "", ErrMissingKID
		}

		var err error
		secret, err = a.GetSecret(kid)
		if err != nil {
			return "", err
		}

		if secret.Status == model.StatusSecretDisabled {
			return "", ErrSecretDisabled
		}

		return []byte(secret.SecretKey), nil
	})

	if err != nil {
		ve, ok := err.(*jwt.ValidationError)
		if !ok {
			return "", errors.Unauthorized(reasonUnauthorized, err.Error())
		}
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return "", jwtauthn.ErrTokenInvalid
		}
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return "", jwtauthn.ErrTokenExpired
		}
		return "", err
	}
	if !token.Valid {
		return "", jwtauthn.ErrTokenInvalid
	}

	if keyExpired(secret.Expires) {
		return "", jwtauthn.ErrTokenExpired
	}
	// you can return claims if you need
	// claims := token.Claims.(*jwt.RegisteredClaims)
	return secret.UserID, nil
}

// GetSecret retrieves the secret using the provided secret ID.
func (a *authnImpl) GetSecret(secretID string) (*model.SecretM, error) {
	// Check if the secret is in the cache.
	if secret, ok := a.secrets.Get(secretID); ok {
		return secret.(*model.SecretM), nil
	}

	// Retrieve the secret from the database.
	secret, err := a.setter.Get(context.Background(), secretID)
	if err != nil {
		return nil, err
	}

	// Add the secret to the cache.
	a.secrets.Add(secretID, secret)
	return secret, nil
}
