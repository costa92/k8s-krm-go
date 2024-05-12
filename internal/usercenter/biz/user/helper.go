package user

import (
	"github.com/costa92/k8s-krm-go/internal/usercenter/model"
	v1 "github.com/costa92/k8s-krm-go/pkg/api/usercenter/v1"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/jinzhu/copier"
)

func ModelToReply(userM *model.UserM) *v1.UserReply {
	var user v1.UserReply
	_ = copier.Copy(&user, userM)
	user.CreatedAt = timestamppb.New(userM.CreatedAt)
	user.UpdatedAt = timestamppb.New(userM.UpdatedAt)
	return &user
}
