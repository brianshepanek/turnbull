package repository

import "github.com/go-redis/redis/v7"

type redisCommentRepository struct {
	redisCommentRepositoryStruct
}

func New(client *redis.Client, namespace string) *redisCommentRepository {
	return &redisCommentRepository{redisCommentRepositoryStruct{
		client:    client,
		namespace: namespace,
	}}
}
