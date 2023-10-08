package post

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/mmvergara/golang-gala/backend/model"
	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	Client *redis.Client
}

func postIDKey(id uuid.UUID) string {
	return fmt.Sprintf("post:%s", id)
}

func (r *RedisRepo) Create(ctx context.Context, post model.Post) error {
	// convert post to json
	postJson, err := json.Marshal(post)
	if err != nil {
		return err
	}

	key := postIDKey(post.PostID)

	txn := r.Client.TxPipeline()

	res := r.Client.SetNX(ctx, key, string(postJson), 0)

	if err := res.Err(); err != nil {
		txn.Discard()
		return fmt.Errorf("failed to set post %w", err)
	}

	if err := r.Client.SAdd(ctx, "posts", key).Err(); err != nil {
		txn.Discard()
		return fmt.Errorf("failed to add post to set %w", err)
	}

	if _, err := txn.Exec(ctx); err != nil {
		return fmt.Errorf("failed to exec transaction %w", err)
	}

	return nil
}

var ErrNotExist = errors.New("post does not exist")

func (r *RedisRepo) FindById(ctx context.Context, id uuid.UUID) (model.Post, error) {
	key := postIDKey(id)

	value, err := r.Client.Get(ctx, key).Result()

	if err != nil {
		fmt.Println("failed to get post")
		return model.Post{}, fmt.Errorf("failed to get Post %w", err)
	}

	if errors.Is(err, redis.Nil) {
		fmt.Printf("id does not exist in the database")
		return model.Post{}, ErrNotExist
	} else if err != nil {
		return model.Post{}, fmt.Errorf("failed to get Post %w", err)
	}

	var post model.Post
	err = json.Unmarshal([]byte(value), &post)

	if err != nil {
		return model.Post{}, fmt.Errorf("failed to unmarshal post %w", err)
	}

	return post, nil
}

func (r *RedisRepo) DeleteByID(ctx context.Context, id uuid.UUID) error {
	key := postIDKey(id)

	txn := r.Client.TxPipeline()

	err := r.Client.Del(ctx, key).Err()

	if errors.Is(err, redis.Nil) {
		txn.Discard()
		return ErrNotExist
	} else if err != nil {
		txn.Discard()
		return fmt.Errorf("failed to delete post %w", err)
	}

	if _, err := txn.Exec(ctx); err != nil {
		return fmt.Errorf("failed to exec %w", err)
	}

	return nil
}

func (r *RedisRepo) Update(ctx context.Context, post model.Post) error {
	// convert post to json
	postJson, err := json.Marshal(post)
	if err != nil {
		return err
	}

	// set post to redis
	key := postIDKey(post.PostID)

	res := r.Client.Set(ctx, key, string(postJson), 0).Err()

	// if post does not exist
	if errors.Is(res, redis.Nil) {
		return ErrNotExist
		// if there is an error
	} else if err != nil {
		return fmt.Errorf("failed to update post %w", err)
	}
	return nil
}

type FindAllPage struct {
	Size   uint64
	Offset uint64
}

type FindResult struct {
	Posts  []model.Post
	Cursor uint64
}

func (r *RedisRepo) FindAll(ctx context.Context, page FindAllPage) (FindResult, error) {
	res := r.Client.SScan(ctx, "posts", page.Offset, "*", int64(page.Size))

	keys, cursor, err := res.Result()

	if err != nil {
		return FindResult{}, fmt.Errorf("failed to get posts %w", err)
	}

	xs, err := r.Client.MGet(ctx, keys...).Result()
	if len(keys) == 0 {
		return FindResult{
			Posts: []model.Post{},
		}, nil
	}

	if err != nil {
		return FindResult{}, fmt.Errorf("failed to get posts %w", err)
	}

	posts := make([]model.Post, len(xs))

	for i, x := range xs {
		x := x.(string)

		var post model.Post

		err := json.Unmarshal([]byte(x), &post)
		if err != nil {
			return FindResult{}, fmt.Errorf("failed to unmarshal post %w", err)
		}

		posts[i] = post

	}

	return FindResult{
		Posts:  posts,
		Cursor: uint64(cursor),
	}, nil

}
