package store

import (
	"context"
	"errors"
	"github.com/costa92/k8s-krm-go/internal/pkg/meta"
	"github.com/costa92/k8s-krm-go/internal/usercenter/model"
	"gorm.io/gorm"
)

type UserStore interface {
	// Create adds a new user record to the database.
	Create(ctx context.Context, user *model.UserM) error
	// List returns a slice of user records based on the specified query conditions.
	List(ctx context.Context, opts ...meta.ListOption) (int64, []*model.UserM, error)
	// Fetch retrieves a user record using provided filters.
	Fetch(ctx context.Context, filters map[string]any) (*model.UserM, error)

	// Get retrieves a user record by userID and username.
	Get(ctx context.Context, userID string, username string) (*model.UserM, error)
	// GetByUsername retrieves a user record using username as the query condition.
	GetByUsername(ctx context.Context, username string) (*model.UserM, error)

	// Update modifies an existing user record.
	Update(ctx context.Context, user *model.UserM) error

	// Delete removes a user record using the provided filters.
	Delete(ctx context.Context, filters map[string]any) error
}

type userStore struct {
	ds *datastore
}

func newUserStore(ds *datastore) *userStore {
	return &userStore{ds: ds}
}

func (d *userStore) db(ctx context.Context) *gorm.DB {
	return d.ds.Core(ctx)
}

// Create adds a new user record to the database.
func (d *userStore) Create(ctx context.Context, user *model.UserM) error {
	return d.db(ctx).Create(&user).Error
}

// List returns a slice of user records based on the specified query conditions.
func (d *userStore) List(ctx context.Context, opts ...meta.ListOption) (int64, []*model.UserM, error) {
	o := meta.NewListOptions(opts...)
	var users []*model.UserM
	db := d.db(ctx)
	if err := db.Limit(o.Limit).Offset(o.Offset).Find(&users).Error; err != nil {
		return 0, nil, err
	}
	var count int64
	if err := db.Model(&model.UserM{}).Count(&count).Error; err != nil {
		return 0, nil, err
	}
	return count, users, nil
}

// Fetch retrieves a user record using provided filters.
func (d *userStore) Fetch(ctx context.Context, filters map[string]any) (*model.UserM, error) {
	var user model.UserM
	if err := d.db(ctx).Where(filters).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Get retrieves a user record by userID and username.
func (d *userStore) Get(ctx context.Context, userID string, username string) (*model.UserM, error) {
	return d.Fetch(ctx, map[string]any{"user_id": userID, "username": username})
}

// GetByUsername retrieves a user record using the provided username.
func (d *userStore) GetByUsername(ctx context.Context, username string) (*model.UserM, error) {
	return d.Fetch(ctx, map[string]any{"username": username})
}

// Update modifies an existing user record.
func (d *userStore) Update(ctx context.Context, user *model.UserM) error {
	return d.db(ctx).Save(user).Error
}

// Delete removes a user record using the provided filters.
func (d *userStore) Delete(ctx context.Context, filters map[string]any) error {
	err := d.db(ctx).Where(filters).Delete(&model.UserM{}).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return nil
}
