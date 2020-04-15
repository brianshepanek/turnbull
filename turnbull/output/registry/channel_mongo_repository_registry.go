package registry

import (
	mongo1 "github.com/brianshepanek/turnbull/turnbull/output/interface/repository/channel/mongo"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

type channelMongoRepositoryRegistry struct {
	client     *mongo.Client
	db         string
	collection string
}

func (r *registry) RegisterMongoChannelRepository(client *mongo.Client, db string, collection string) {
	r.channelMongoRepositoryRegistry.client = client
	r.channelMongoRepositoryRegistry.db = db
	r.channelMongoRepositoryRegistry.collection = collection
}
func (r *registry) newMongoChannelRepository() repository.ChannelRepository {
	return mongo1.New(r.channelMongoRepositoryRegistry.client, r.channelMongoRepositoryRegistry.db, r.channelMongoRepositoryRegistry.collection)
}
