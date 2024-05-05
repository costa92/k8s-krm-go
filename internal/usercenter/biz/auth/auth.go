package auth

import (
	"context"
	"github.com/costa92/k8s-krm-go/internal/pkg/krmx"
	"github.com/costa92/k8s-krm-go/internal/usercenter/auth"
	"github.com/costa92/k8s-krm-go/internal/usercenter/locales"
	"github.com/costa92/k8s-krm-go/internal/usercenter/store"
	v1 "github.com/costa92/k8s-krm-go/pkg/api/usercenter/v1"
	"github.com/costa92/k8s-krm-go/pkg/authn"
	"github.com/costa92/k8s-krm-go/pkg/i18n"
	"github.com/costa92/k8s-krm-go/pkg/log"
)

type AuthBiz interface {
	// Login authenticates a user and returns a token.
	Login(ctx context.Context, rq *v1.LoginRequest) (*v1.LoginReply, error)

	// Logout invalidates a token.
	Logout(ctx context.Context, rq *v1.LogoutRequest) error
}

type authBiz struct {
	ds    store.IStore
	authn authn.Authenticator
	auth  auth.AuthProvider
}

var _ AuthBiz = (*authBiz)(nil)

func New(ds store.IStore, authn authn.Authenticator, auth auth.AuthProvider) *authBiz {
	return &authBiz{
		ds:    ds,
		authn: authn,
		auth:  auth,
	}
}

// Login authenticates a user and returns a token.
func (b *authBiz) Login(ctx context.Context, rq *v1.LoginRequest) (*v1.LoginReply, error) {
	// Retrieve user information from the data storage by username.
	userM, err := b.ds.Users().GetByUsername(ctx, rq.Username)
	if err != nil {
		log.C(ctx).Errorw(err, "Failed to retrieve user by username")
		return nil, i18n.FromContext(ctx).E(locales.RecordNotFound)
	}
	// Compare the password.
	if err := authn.Compare(userM.Password, rq.Password); err != nil {
		log.C(ctx).Errorw(err, "Failed to compare password")
		return nil, i18n.FromContext(ctx).E(locales.IncorrectPassword)
	}
	// Generate a token.
	refreshToken, err := b.authn.Sign(ctx, userM.UserID)
	if err != nil {
		log.C(ctx).Errorw(err, "Failed to generate refresh token")
		return nil, i18n.FromContext(ctx).E(locales.JWTTokenSignFail)
	}

	accessToken, err := b.auth.Sign(ctx, userM.UserID)
	if err != nil {
		log.C(ctx).Errorw(err, "Failed to generate access token")
		return nil, i18n.FromContext(ctx).E(locales.JWTTokenSignFail)
	}
	// Return
	return &v1.LoginReply{
		RefreshToken: refreshToken.GetToken(),
		AccessToken:  accessToken.GetToken(),
		Type:         accessToken.GetTokenType(),
		ExpiresAt:    accessToken.GetExpiresAt(),
	}, nil
}

// Logout invalidates a token.
func (b *authBiz) Logout(ctx context.Context, rq *v1.LogoutRequest) error {
	if err := b.authn.Destroy(ctx, krmx.FromAccessToken(ctx)); err != nil {
		log.C(ctx).Errorw(err, "Failed to remove token from cache")
		return err
	}

	return nil
}
