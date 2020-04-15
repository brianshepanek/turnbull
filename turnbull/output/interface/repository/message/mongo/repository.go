package repository

import mongo "go.mongodb.org/mongo-driver/mongo"

type mongoMessageRepository struct {
	mongoMessageRepositoryStruct
}

func New(client *mongo.Client, db string, collection string) *mongoMessageRepository {
	return &mongoMessageRepository{mongoMessageRepositoryStruct{
		client:     client,
		collection: collection,
		db:         db,
	}}
}
