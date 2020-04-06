package repository

import "github.com/go-redis/redis/v7"

type redisPostRepository struct {
	redisPostRepositoryStruct
}

func New(client *redis.Client, namespace string) *redisPostRepository {
	return &redisPostRepository{redisPostRepositoryStruct{
		client:    client,
		namespace: namespace,
	}}
}
