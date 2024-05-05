package jwt

import (
	"context"
	"time"
)

type Storer interface {
	// Set stores the token data and specifies the expiration time.
	Set(ctx context.Context, accessToken string, expiration time.Duration) error
	// Delete deletes the token data from the storage.
	Delete(ctx context.Context, accessToken string) (bool, error)
	// Check checks if the token exists.
	Check(ctx context.Context, accessToken string) (bool, error)
	// Close closes the storage.
	Close() error
}
