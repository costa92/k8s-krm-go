package usercenter

import "time"

const (
	DefaultLRUSize = 1000
	// AccessTokenExpire is the expiration time for the access token.
	AccessTokenExpire = time.Hour * 2
	// RefreshTokenExpire is the expiration time for the refresh token.
	RefreshTokenExpire = time.Hour * 24
)

const (
	TemporaryKeyName = "_krm.io/temporary_key"
)
