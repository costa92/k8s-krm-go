package validation

import (
	"context"
	"fmt"
	"github.com/costa92/k8s-krm-go/internal/pkg/known"
	"github.com/costa92/k8s-krm-go/internal/pkg/known/usercenter"
	"github.com/costa92/k8s-krm-go/internal/pkg/krmx"
	"github.com/costa92/k8s-krm-go/internal/usercenter/locales"
	"github.com/costa92/k8s-krm-go/internal/usercenter/store"
	v1 "github.com/costa92/k8s-krm-go/pkg/api/usercenter/v1"
	"github.com/costa92/k8s-krm-go/pkg/i18n"
)

type validator struct {
	ds store.IStore
}

func New(ds store.IStore) (*validator, error) {
	return &validator{ds: ds}, nil
}

func (vd *validator) ValidateCreateUserRequest(ctx context.Context, rq *v1.CreateUserRequest) error {
	if _, err := vd.ds.Users().GetByUsername(ctx, rq.Username); err == nil {
		return i18n.FromContext(ctx).E(locales.UserAlreadyExists)
	}
	return nil

}

func (vd *validator) ValidateListUserRequest(ctx context.Context, rq *v1.ListUserRequest) error {
	if userID := krmx.FromUserID(ctx); userID == known.AdminUserID {
		return i18n.FromContext(ctx).E(locales.UserListUnauthorized)
	}
	return nil

}

func (vd *validator) ValidateCreateSecretRequest(ctx context.Context, rq *v1.GetUserRequest) error {
	_, secrets, err := vd.ds.Secret().List(ctx, krmx.FromUserID(ctx))
	if err != nil {
		return err
	}
	if len(secrets) >= usercenter.MaxSecretCount {
		return fmt.Errorf("secret reach the max count %d", usercenter.MaxSecretCount)
	}
	return nil
}
