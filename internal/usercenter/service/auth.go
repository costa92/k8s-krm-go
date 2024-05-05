package service

import (
	"context"
	v1 "github.com/costa92/k8s-krm-go/pkg/api/usercenter/v1"
)

func (s *UserCenterService) Login(ctx context.Context, rq *v1.LoginRequest) (*v1.LoginReply, error) {
	resp, err := s.biz.Auths().Login(ctx, rq)
	if err != nil {
		return &v1.LoginReply{}, v1.ErrorUserLoginFailed(err.Error())
	}
	return resp, nil
}
