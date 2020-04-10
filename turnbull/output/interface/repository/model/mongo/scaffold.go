package repository

import mongo "go.mongodb.org/mongo-driver/mongo"

type mongoModelRepositoryStruct struct {
	client     *mongo.Client
	db         string
	collection string
}
