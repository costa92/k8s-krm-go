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

func (s *UserCenterService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.UserReply, error) {
	log.C(ctx).Infow("GetUser Function called", "username", req.Username)
	return s.biz.Users().Get(ctx, req)

}

func (s *UserCenterService) ListUser(ctx context.Context, req *v1.ListUserRequest) (*v1.ListUserResponse, error) {
	return s.biz.Users().List(ctx, req)
}
