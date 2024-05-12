package user

import (
	"context"
	"github.com/costa92/k8s-krm-go/internal/usercenter/store"
	v1 "github.com/costa92/k8s-krm-go/pkg/api/usercenter/v1"
)

type UserBiz interface {
	Create(ctx context.Context, user *v1.CreateUserRequest) (*v1.UserReply, error)
}

type userBiz struct {
	ds store.IStore
}

var _ UserBiz = (*userBiz)(nil)

func New(ds store.IStore) *userBiz {
	return &userBiz{ds: ds}
}
