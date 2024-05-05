package store

import (
	"context"
	known "github.com/costa92/k8s-krm-go/internal/pkg/known/usercenter"
	"github.com/costa92/k8s-krm-go/internal/usercenter/auth"
	"github.com/costa92/k8s-krm-go/internal/usercenter/model"
	"github.com/google/wire"
)

type secretSetter struct {
	ds *datastore
}

var (
	SetterProviderSet                            = wire.NewSet(NewSecretSetter, wire.Bind(new(auth.TemporarySecretSetter), new(*secretSetter)))
	_                 auth.TemporarySecretSetter = (*secretSetter)(nil)
)

// NewSecretSetter initializes a new secretSetter instance using the provided datastore.
func NewSecretSetter(ds *datastore) *secretSetter {
	return &secretSetter{ds}
}

// Get retrieves a secret record from the datastore based on userID and secretID.
func (d *secretSetter) Get(ctx context.Context, secretID string) (*model.SecretM, error) {
	secret := &model.SecretM{}
	if err := d.ds.core.Where(model.SecretM{SecretID: secretID}).First(&secret).Error; err != nil {
		return nil, err
	}

	return secret, nil
}

// Create adds a new secret record in the datastore.
func (d *secretSetter) Set(ctx context.Context, userID string, expires int64) (*model.SecretM, error) {
	var secret model.SecretM
	err := d.ds.core.
		Where(model.SecretM{Name: known.TemporaryKeyName, UserID: userID}).
		Assign(model.SecretM{Expires: expires}).
		FirstOrCreate(&secret).
		Error
	if err != nil {
		return nil, err
	}

	return &secret, nil
}
