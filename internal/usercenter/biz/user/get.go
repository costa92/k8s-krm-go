package user

import (
	"context"
	v1 "github.com/costa92/k8s-krm-go/pkg/api/usercenter/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
)

func (b *userBiz) Get(ctx context.Context, req *v1.GetUserRequest) (*v1.UserReply, error) {
	filters := map[string]interface{}{
		"username": req.Username,
	}
	userM, err := b.ds.Users().Fetch(ctx, filters)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrorUserNotFound(err.Error())
		}
		return nil, err
	}
	return ModelToReply(userM), nil
}
