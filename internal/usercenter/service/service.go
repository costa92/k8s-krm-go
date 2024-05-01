package service

import (
	v1 "github.com/costa92/k8s-krm-go/pkg/api/usercenter/v1"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewUserCenterService)

type UserCenterService struct {
	v1.UnimplementedUserServiceServer
}

func NewUserCenterService() *UserCenterService {
	return &UserCenterService{}
}
