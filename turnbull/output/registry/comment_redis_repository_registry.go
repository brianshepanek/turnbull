package registry

import (
	redis1 "github.com/brianshepanek/turnbull/turnbull/output/interface/repository/comment/redis"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
	"github.com/go-redis/redis/v7"
)

type commentRedisRepositoryRegistry struct {
	client    *redis.Client
	namespace string
}

func (r *registry) RegisterCommentRedisRepositoryRegistry(client *redis.Client, namespace string) {
	r.commentRedisRepositoryRegistry.client = client
	r.commentRedisRepositoryRegistry.namespace = namespace
}
func (r *registry) newCommentRedisRepositoryRegistry() repository.CommentRepository {
	return redis1.New(r.commentRedisRepositoryRegistry.client, r.commentRedisRepositoryRegistry.namespace)
}
