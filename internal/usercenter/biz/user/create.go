package user

import (
	"context"
	"github.com/costa92/k8s-krm-go/internal/usercenter/model"
	v1 "github.com/costa92/k8s-krm-go/pkg/api/usercenter/v1"
	"github.com/jinzhu/copier"
	"regexp"
	"time"
)

func (b *userBiz) Create(ctx context.Context, req *v1.CreateUserRequest) (*v1.UserReply, error) {
	var userM model.UserM
	_ = copier.Copy(&userM, req)
	userM.CreatedAt = time.Now()
	userM.UpdatedAt = time.Now()
	err := b.ds.TX(ctx, func(ctx context.Context) error {
		if err := b.ds.Users().Create(ctx, &userM); err != nil {
			match, _ := regexp.MatchString("Duplicate entry '.*' for key 'username'", err.Error())
			if match {
				return v1.ErrorUserAlreadyExists("usr %q already exists", userM.Username)
			}
			return v1.ErrorUserCreateFailed("failed to create user: %v", err.Error())
		}

		secretM := &model.SecretM{
			UserID:      userM.UserID,
			Name:        "generated",
			Expires:     0,
			Description: "automatically generated when user is created",
		}
		if err := b.ds.Secret().Create(ctx, secretM); err != nil {
			return v1.ErrorSecretCreateFailed("failed to create secret: %v", err.Error())
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return ModelToReply(&userM), nil
}
