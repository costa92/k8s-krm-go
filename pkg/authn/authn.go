package authn

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IToken interface {
	// GetToken is a method that returns the token.
	GetToken() string
	// GetTokenType is a method that returns the token type.
	GetTokenType() string
	// GetExpiresAt is a method that returns the expiration time of the token.
	GetExpiresAt() int64
	// EncodeToJSON is a method that encodes the token to JSON.
	EncodeToJSON() ([]byte, error)
}

type Authenticator interface {
	// Sign is a method that generates a token.
	Sign(ctx context.Context, userID string) (IToken, error)
	// Destroy is a method that destroys a token.
	Destroy(ctx context.Context, accessToken string) error
	// ParseClaims is a method that parses the token and returns the claims.
	ParseClaims(ctx context.Context, accessToken string) (*jwt.RegisteredClaims, error)
	// Release is a method that releases the requested resources.
	Release() error
}

// Encrypt encrypts the plain text with bcrypt.
func Encrypt(source string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

// Compare compares the encrypted text with the plain text if it's the same.
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
