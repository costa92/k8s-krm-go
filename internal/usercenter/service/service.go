package service

import (
	"github.com/costa92/k8s-krm-go/internal/usercenter/biz"
	v1 "github.com/costa92/k8s-krm-go/pkg/api/usercenter/v1"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewUserCenterService)

type UserCenterService struct {
	v1.UnimplementedUserServiceServer
	biz biz.IBiz
}

func NewUserCenterService(biz biz.IBiz) *UserCenterService {
	return &UserCenterService{
		biz: biz,
	}
}
