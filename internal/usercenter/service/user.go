package service

import (
	"context"
	v1 "github.com/costa92/k8s-krm-go/pkg/api/usercenter/v1"
	"github.com/costa92/k8s-krm-go/pkg/log"
)

func (s *UserCenterService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.UserReply, error) {
	log.C(ctx).Infow("CreateUser Function called", "username", req.Username)
	return s.biz.Users().Create(ctx, req)
}
