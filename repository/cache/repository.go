package cache

import (
	"context"

	"github.com/BRO3886/url-shortener/shortener"
	"github.com/go-redis/redis/v8"
)

type redisRepo struct {
	client *redis.Client
}

// Create implements shortener.RedirectRepository
func (r *redisRepo) Create(redirect *shortener.Redirect) (*shortener.Redirect, error) {
	key := r.genKey(redirect.Code)
	data := redirect.ToMap()

	_, err := r.client.HSet(context.Background(), key, data).Result()
	if err != nil {
		return nil, err
	}

	return redirect, nil
}

// Find implements shortener.RedirectRepository
func (r *redisRepo) Find(code string) (*shortener.Redirect, error) {
	redirect := new(shortener.Redirect)
	key := r.genKey(code)

	data, err := r.client.HGetAll(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, shortener.ErrRedirectNotFound
	}

	redirect.FromMap(data)

	return redirect, nil
}

func NewRepository(redisURL, username, password string) (shortener.RedirectRepository, error) {
	client, err := newRedisClient(redisURL, username, password)
	if err != nil {
		return nil, err
	}
	return &redisRepo{
		client: client,
	}, nil
}

func (*redisRepo) genKey(code string) string {
	return "redirect:" + code
}
