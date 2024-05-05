package biz

//go:generate mockgen -self_package github.com/costa92/k8s-krm-go/internal/usercenter/biz -destination mock_biz.go -package biz github.com/costa92/k8s-krm-go/internal/usercenter/biz IBiz

import (
	"github.com/costa92/k8s-krm-go/internal/usercenter/store"
	"github.com/google/wire"
)

// ProviderSet contains providers for creating instances of the biz struct.
var ProviderSet = wire.NewSet(NewBiz, wire.Bind(new(IBiz), new(*biz)))

type IBiz interface {
}

type biz struct {
	ds store.IStore
}

func NewBiz(ds store.IStore) *biz {
	return &biz{ds: ds}
}
