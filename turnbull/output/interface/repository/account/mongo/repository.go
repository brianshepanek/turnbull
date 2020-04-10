package repository

import mongo "go.mongodb.org/mongo-driver/mongo"

type mongoAccountRepository struct {
	mongoAccountRepositoryStruct
}

func New(client *mongo.Client, db string, collection string) *mongoAccountRepository {
	return &mongoAccountRepository{mongoAccountRepositoryStruct{
		client:     client,
		collection: collection,
		db:         db,
	}}
}
