package user

import (
	"context"
	"github.com/costa92/k8s-krm-go/internal/usercenter/store"
	v1 "github.com/costa92/k8s-krm-go/pkg/api/usercenter/v1"
)

type UserBiz interface {
	// Create creates a new user.
	Create(ctx context.Context, user *v1.CreateUserRequest) (*v1.UserReply, error)
	// Get returns the user by username.
	Get(ctx context.Context, req *v1.GetUserRequest) (*v1.UserReply, error)
	// List returns a list of users.
	List(ctx context.Context, req *v1.ListUserRequest) (*v1.ListUserResponse, error)
}

type userBiz struct {
	ds store.IStore
}

var _ UserBiz = (*userBiz)(nil)

func New(ds store.IStore) *userBiz {
	return &userBiz{ds: ds}
}
