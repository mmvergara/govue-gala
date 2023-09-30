package user

import (
	"context"
	"encoding/json"

	"github.com/mmvergara/golang-gala/backend/model"
	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	Client *redis.Client
}


func (r *RedisRepo) Insert(ctx context.Context, user model.User) error {
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}

	if err := r.Client.Set(ctx, user.UserID.String(), userJSON, 0).Err(); err != nil {
		return err
	}

	return nil
}