package user

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/mmvergara/golang-gala/backend/model"
	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	Client *redis.Client
}


func (r *RedisRepo) CreateUser(ctx context.Context, user model.User) error {
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}

	if err := r.Client.Set(ctx, user.UserID.String(), userJSON, 0).Err(); err != nil {
		return err
	}

	return nil
}

func (r *RedisRepo) FindById(ctx context.Context,userID uuid.UUID) (model.User, error) {
	userJSON, err := r.Client.Get(ctx, userID.String()).Result()
	if err != nil {
		return model.User{}, err
	}

	var user model.User
	if err := json.Unmarshal([]byte(userJSON), &user); err != nil {
		return model.User{}, err
	}

	return user, nil
}