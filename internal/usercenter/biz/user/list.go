package user

import (
	"context"
	"github.com/costa92/k8s-krm-go/internal/pkg/meta"
	v1 "github.com/costa92/k8s-krm-go/pkg/api/usercenter/v1"
	"github.com/costa92/k8s-krm-go/pkg/log"
	"github.com/jinzhu/copier"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/types/known/timestamppb"
	"sync"
)

func (b *userBiz) List(ctx context.Context, req *v1.ListUserRequest) (*v1.ListUserResponse, error) {
	count, list, err := b.ds.Users().List(ctx, meta.WithOffset(req.Offset), meta.WithLimit(req.Limit))
	if err != nil {
		log.C(ctx).Errorw(err, "failed to list users")
		return nil, err
	}

	var m sync.Map
	eg, ctx := errgroup.WithContext(ctx)
	// fetch user secrets
	for _, user := range list {
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return nil
			default:
				count, _, err := b.ds.Secret().List(ctx, user.UserID)
				if err != nil {
					log.C(ctx).Errorw(err, "failed to list secrets")
					return err
				}
				u := ModelToReply(user)
				u.Secrets = count
				m.Store(user.UserID, u)
				return nil
			}
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	users := make([]*v1.UserReply, 0, len(list))
	for _, item := range list {
		if v, ok := m.Load(item.UserID); ok {
			users = append(users, v.(*v1.UserReply))
		}
	}
	log.C(ctx).Debugw("list users", "count", len(users))

	return &v1.ListUserResponse{
		TotalCount: count,
		Users:      users,
	}, nil
}

func (b *userBiz) ListWithBadPerformance(ctx context.Context, req *v1.ListUserRequest) (*v1.ListUserResponse, error) {
	count, list, err := b.ds.Users().List(ctx, meta.WithOffset(req.Offset), meta.WithLimit(req.Limit))
	if err != nil {
		log.C(ctx).Errorw(err, "Failed to list users from storage")
		return nil, err
	}

	users := make([]*v1.UserReply, 0)
	for _, user := range list {
		var u v1.UserReply
		_ = copier.Copy(&u, user)
		secretCount, _, err := b.ds.Secret().List(ctx, user.UserID)
		if err != nil {
			log.C(ctx).Errorw(err, "Failed to list secrets from storage")
			return nil, err
		}
		u.CreatedAt = timestamppb.New(user.CreatedAt)
		u.UpdatedAt = timestamppb.New(user.UpdatedAt)
		u.Password = "******"
		u.Secrets = secretCount
		users = append(users, &u)
	}
	log.C(ctx).Debugw("List users", "count", len(users))
	return &v1.ListUserResponse{
		TotalCount: count,
		Users:      users,
	}, nil
}
