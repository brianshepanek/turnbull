package repository

import mongo "go.mongodb.org/mongo-driver/mongo"

type mongoEnhancedAccountRepository struct {
	mongoEnhancedAccountRepositoryStruct
}

func New(client *mongo.Client, db string, collection string) *mongoEnhancedAccountRepository {
	return &mongoEnhancedAccountRepository{mongoEnhancedAccountRepositoryStruct{
		client:     client,
		collection: collection,
		db:         db,
	}}
}
