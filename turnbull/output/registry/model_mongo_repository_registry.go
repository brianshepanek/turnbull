package registry

import (
	mongo1 "github.com/brianshepanek/turnbull/turnbull/output/interface/repository/model/mongo"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

type modelMongoRepositoryRegistry struct {
	client     *mongo.Client
	db         string
	collection string
}

func (r *registry) RegisterMongoModelRepository(client *mongo.Client, db string, collection string) {
	r.modelMongoRepositoryRegistry.client = client
	r.modelMongoRepositoryRegistry.db = db
	r.modelMongoRepositoryRegistry.collection = collection
}
func (r *registry) newMongoModelRepository() repository.ModelRepository {
	return mongo1.New(r.modelMongoRepositoryRegistry.client, r.modelMongoRepositoryRegistry.db, r.modelMongoRepositoryRegistry.collection)
}
