package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func newRedisClient(redisURL, username, password string) (*redis.Client, error) {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Username: username,
		Password: password,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
