package repository

import mongo "go.mongodb.org/mongo-driver/mongo"

type mongoFooRepository struct {
	mongoFooRepositoryStruct
}

func New(client *mongo.Client, db string, collection string) *mongoFooRepository {
	return &mongoFooRepository{mongoFooRepositoryStruct{
		client:     client,
		collection: collection,
		db:         db,
	}}
}
