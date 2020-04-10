package repository

import mongo "go.mongodb.org/mongo-driver/mongo"

type mongoModelRepository struct {
	mongoModelRepositoryStruct
}

func New(client *mongo.Client, db string, collection string) *mongoModelRepository {
	return &mongoModelRepository{mongoModelRepositoryStruct{
		client:     client,
		collection: collection,
		db:         db,
	}}
}
