package repository

import mongo "go.mongodb.org/mongo-driver/mongo"

type mongoPostRepository struct {
	mongoPostRepositoryStruct
}

func New(client *mongo.Client, db string, collection string) *mongoPostRepository {
	return &mongoPostRepository{mongoPostRepositoryStruct{
		client:     client,
		collection: collection,
		db:         db,
	}}
}
