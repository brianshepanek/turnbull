package repository

import mongo "go.mongodb.org/mongo-driver/mongo"

type mongoUserRepository struct {
	mongoUserRepositoryStruct
}

func New(client *mongo.Client, db string, collection string) *mongoUserRepository {
	return &mongoUserRepository{mongoUserRepositoryStruct{
		client:     client,
		collection: collection,
		db:         db,
	}}
}
