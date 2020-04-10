package registry

import (
	mongo1 "github.com/brianshepanek/turnbull/turnbull/output/interface/repository/account/mongo"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

type accountMongoRepositoryRegistry struct {
	client     *mongo.Client
	db         string
	collection string
}

func (r *registry) RegisterMongoAccountRepository(client *mongo.Client, db string, collection string) {
	r.accountMongoRepositoryRegistry.client = client
	r.accountMongoRepositoryRegistry.db = db
	r.accountMongoRepositoryRegistry.collection = collection
}
func (r *registry) newMongoAccountRepository() repository.AccountRepository {
	return mongo1.New(r.accountMongoRepositoryRegistry.client, r.accountMongoRepositoryRegistry.db, r.accountMongoRepositoryRegistry.collection)
}
