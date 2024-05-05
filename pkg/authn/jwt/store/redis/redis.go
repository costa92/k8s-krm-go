package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type Config struct {
	Addr     string
	Username string
	Password string
	Database int
	// Sore key prefix.
	KeyPrefix string
}

type Store struct {
	cli    *redis.Client
	prefix string
}

func NewStore(cfg *Config) *Store {
	cli := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		DB:       cfg.Database,
		Username: cfg.Username,
		Password: cfg.Password,
	})
	return &Store{cli: cli, prefix: cfg.KeyPrefix}
}

func (s *Store) wrapperKey(key string) string {
	return fmt.Sprintf("%s%s", s.prefix, key)
}

func (s *Store) Set(ctx context.Context, accessToken string, expiration time.Duration) error {
	cmd := s.cli.Set(ctx, s.wrapperKey(accessToken), "1", expiration)
	return cmd.Err()
}

// Delete delete the specified JWT Token in Redis.
func (s *Store) Delete(ctx context.Context, accessToken string) (bool, error) {
	cmd := s.cli.Del(ctx, s.wrapperKey(accessToken))
	if err := cmd.Err(); err != nil {
		return false, err
	}
	return cmd.Val() > 0, nil
}

// Check if the specified JWT Token exists in Redis.
func (s *Store) Check(ctx context.Context, accessToken string) (bool, error) {
	cmd := s.cli.Exists(ctx, s.wrapperKey(accessToken))
	if err := cmd.Err(); err != nil {
		return false, err
	}
	return cmd.Val() > 0, nil
}

// Close is used to close the redis client.
func (s *Store) Close() error {
	return s.cli.Close()
}
