package repository

import mongo "go.mongodb.org/mongo-driver/mongo"

type mongoCommentRepository struct {
	mongoCommentRepositoryStruct
}

func New(client *mongo.Client, db string, collection string) *mongoCommentRepository {
	return &mongoCommentRepository{mongoCommentRepositoryStruct{
		client:     client,
		collection: collection,
		db:         db,
	}}
}
