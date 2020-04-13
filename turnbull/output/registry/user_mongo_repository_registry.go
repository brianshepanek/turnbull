package registry

import (
	mongo1 "github.com/brianshepanek/turnbull/turnbull/output/interface/repository/user/mongo"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

type userMongoRepositoryRegistry struct {
	client     *mongo.Client
	db         string
	collection string
}

func (r *registry) RegisterMongoUserRepository(client *mongo.Client, db string, collection string) {
	r.userMongoRepositoryRegistry.client = client
	r.userMongoRepositoryRegistry.db = db
	r.userMongoRepositoryRegistry.collection = collection
}
func (r *registry) newMongoUserRepository() repository.UserRepository {
	return mongo1.New(r.userMongoRepositoryRegistry.client, r.userMongoRepositoryRegistry.db, r.userMongoRepositoryRegistry.collection)
}
