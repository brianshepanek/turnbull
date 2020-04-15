package repository

import mongo "go.mongodb.org/mongo-driver/mongo"

type mongoChannelRepository struct {
	mongoChannelRepositoryStruct
}

func New(client *mongo.Client, db string, collection string) *mongoChannelRepository {
	return &mongoChannelRepository{mongoChannelRepositoryStruct{
		client:     client,
		collection: collection,
		db:         db,
	}}
}
