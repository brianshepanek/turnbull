package registry

import (
	mongo1 "github.com/brianshepanek/turnbull/turnbull/output/interface/repository/comment/mongo"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

type commentMongoRepositoryRegistry struct {
	client     *mongo.Client
	db         string
	collection string
}

func (r *registry) RegisterMongoCommentRepository(client *mongo.Client, db string, collection string) {
	r.commentMongoRepositoryRegistry.client = client
	r.commentMongoRepositoryRegistry.db = db
	r.commentMongoRepositoryRegistry.collection = collection
}
func (r *registry) newMongoCommentRepository() repository.CommentRepository {
	return mongo1.New(r.commentMongoRepositoryRegistry.client, r.commentMongoRepositoryRegistry.db, r.commentMongoRepositoryRegistry.collection)
}
