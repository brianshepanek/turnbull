package registry

import (
	mongo1 "github.com/brianshepanek/turnbull/turnbull/output/interface/repository/post/mongo"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

type postMongoRepositoryRegistry struct {
	client     *mongo.Client
	db         string
	collection string
}

func (r *registry) RegisterPostMongoRepositoryRegistry(client *mongo.Client, db string, collection string) {
	r.postMongoRepositoryRegistry.client = client
	r.postMongoRepositoryRegistry.db = db
	r.postMongoRepositoryRegistry.collection = collection
}
func (r *registry) newPostMongoRepository() repository.PostRepository {
	return mongo1.New(r.postMongoRepositoryRegistry.client, r.postMongoRepositoryRegistry.db, r.postMongoRepositoryRegistry.collection)
}
