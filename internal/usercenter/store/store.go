package store

import (
	"context"
	"github.com/google/wire"
	"gorm.io/gorm"
	"sync"
)

var ProviderSet = wire.NewSet(NewStore, wire.Bind(new(IStore), new(*datastore)))

var (
	once sync.Once
	S    *datastore
)

type transactionKey struct{}

type IStore interface {
	TX(ctx context.Context, fn func(ctx context.Context) error) error
	// Users returns a UserStore on which user operations can be performed.
	Users() UserStore
}

type datastore struct {
	core *gorm.DB
}

var _ IStore = (*datastore)(nil)

func NewStore(db *gorm.DB) *datastore {
	once.Do(func() {
		S = &datastore{db}
	})
	return S
}

func (ds *datastore) Core(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(transactionKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return ds.core
}

func (ds *datastore) TX(ctx context.Context, fn func(ctx context.Context) error) error {
	return ds.core.WithContext(ctx).Transaction(
		func(tx *gorm.DB) error {
			ctx = context.WithValue(ctx, transactionKey{}, tx)
			return fn(ctx)
		},
	)
}

// Users returns a UserStore on which user operations can be performed.
func (ds *datastore) Users() UserStore {
	return newUserStore(ds)
}
