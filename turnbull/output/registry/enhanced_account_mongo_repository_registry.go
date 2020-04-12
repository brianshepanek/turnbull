package registry

import (
	mongo1 "github.com/brianshepanek/turnbull/turnbull/output/interface/repository/enhanced_account/mongo"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

type enhancedAccountMongoRepositoryRegistry struct {
	client     *mongo.Client
	db         string
	collection string
}

func (r *registry) RegisterMongoEnhancedAccountRepository(client *mongo.Client, db string, collection string) {
	r.enhancedAccountMongoRepositoryRegistry.client = client
	r.enhancedAccountMongoRepositoryRegistry.db = db
	r.enhancedAccountMongoRepositoryRegistry.collection = collection
}
func (r *registry) newMongoEnhancedAccountRepository() repository.EnhancedAccountRepository {
	return mongo1.New(r.enhancedAccountMongoRepositoryRegistry.client, r.enhancedAccountMongoRepositoryRegistry.db, r.enhancedAccountMongoRepositoryRegistry.collection)
}
