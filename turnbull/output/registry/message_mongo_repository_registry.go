package registry

import (
	mongo1 "github.com/brianshepanek/turnbull/turnbull/output/interface/repository/message/mongo"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

type messageMongoRepositoryRegistry struct {
	client     *mongo.Client
	db         string
	collection string
}

func (r *registry) RegisterMongoMessageRepository(client *mongo.Client, db string, collection string) {
	r.messageMongoRepositoryRegistry.client = client
	r.messageMongoRepositoryRegistry.db = db
	r.messageMongoRepositoryRegistry.collection = collection
}
func (r *registry) newMongoMessageRepository() repository.MessageRepository {
	return mongo1.New(r.messageMongoRepositoryRegistry.client, r.messageMongoRepositoryRegistry.db, r.messageMongoRepositoryRegistry.collection)
}
