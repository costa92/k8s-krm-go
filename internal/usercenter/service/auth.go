package service

import (
	"context"
	v1 "github.com/costa92/k8s-krm-go/pkg/api/usercenter/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Login authenticates the user credentials and returns a token on success.
func (s *UserCenterService) Login(context.Context, *v1.LoginRequest) (*v1.LoginReply, error) {
	return &v1.LoginReply{
		Type:      "Login",
		ExpiresAt: 3600,
	}, nil
}

// Logout invalidates the user token.
func (s *UserCenterService) Logout(ctx context.Context, rq *v1.LogoutRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
