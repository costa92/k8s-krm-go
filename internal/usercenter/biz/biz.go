package biz

//go:generate mockgen -self_package github.com/costa92/k8s-krm-go/internal/usercenter/biz -destination mock_biz.go -package biz github.com/costa92/k8s-krm-go/internal/usercenter/biz IBiz

import (
	"github.com/costa92/k8s-krm-go/internal/usercenter/auth"
	authbiz "github.com/costa92/k8s-krm-go/internal/usercenter/biz/auth"
	"github.com/costa92/k8s-krm-go/internal/usercenter/store"
	"github.com/costa92/k8s-krm-go/pkg/authn"
	"github.com/google/wire"
)

// ProviderSet contains providers for creating instances of the biz struct.
var ProviderSet = wire.NewSet(NewBiz, wire.Bind(new(IBiz), new(*biz)))

type IBiz interface {
	Auths() authbiz.AuthBiz
}

type biz struct {
	ds    store.IStore
	authn authn.Authenticator
	auth  auth.AuthProvider
}

func NewBiz(ds store.IStore, authn authn.Authenticator, auth auth.AuthProvider) *biz {
	return &biz{ds: ds, authn: authn, auth: auth}
}

// Auths returns a new instance of the AuthBiz interface.
func (b *biz) Auths() authbiz.AuthBiz {
	return authbiz.New(b.ds, b.authn, b.auth)
}
