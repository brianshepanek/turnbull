package registry

import (
	redis1 "github.com/brianshepanek/turnbull/turnbull/output/interface/repository/post/redis"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
	"github.com/go-redis/redis/v7"
)

type postRedisRepositoryRegistry struct {
	client    *redis.Client
	namespace string
}

func (r *registry) RegisterRedisPostRepository(client *redis.Client, namespace string) {
	r.postRedisRepositoryRegistry.client = client
	r.postRedisRepositoryRegistry.namespace = namespace
}
func (r *registry) newRedisPostRepository() repository.PostRepository {
	return redis1.New(r.postRedisRepositoryRegistry.client, r.postRedisRepositoryRegistry.namespace)
}
