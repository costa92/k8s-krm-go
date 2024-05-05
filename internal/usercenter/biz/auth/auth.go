package auth

import (
	"context"
	"github.com/costa92/k8s-krm-go/internal/usercenter/store"
	v1 "github.com/costa92/k8s-krm-go/pkg/api/usercenter/v1"
	"github.com/costa92/k8s-krm-go/pkg/log"
)

type AuthBiz interface {
	// Login authenticates a user and returns a token.
	Login(ctx context.Context, rq *v1.LoginRequest) (*v1.LoginReply, error)

	// Logout invalidates a token.
	Logout(ctx context.Context, rq *v1.LogoutRequest) error
}

type authBiz struct {
	ds store.IStore
}

var _ AuthBiz = (*authBiz)(nil)

func New(ds store.IStore) *authBiz {
	return &authBiz{ds: ds}
}

// Login authenticates a user and returns a token.
func (b *authBiz) Login(ctx context.Context, rq *v1.LoginRequest) (*v1.LoginReply, error) {
	// Retrieve user information from the data storage by username.
	userM, err := b.ds.Users().GetByUsername(ctx, rq.Username)
	if err != nil {
		log.C(ctx).Errorw(err, "Failed to retrieve user by username")
		return nil, err
	}
	return nil, nil
}

// Logout invalidates a token.
func (b *authBiz) Logout(ctx context.Context, rq *v1.LogoutRequest) error {
	return nil
}
