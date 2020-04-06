package repository

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
	bson "go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

type mongoPostRepositoryStruct struct {
	client     *mongo.Client
	db         string
	collection string
}

func (r *mongoPostRepositoryStruct) Browse(ctx context.Context, req *[]entity.Post) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return err
	}

	for cursor.Next(ctx) {
		elem := entity.NewPost()
		err := cursor.Decode(&elem)
		if err != nil {
			return err
		}
		*req = append(*req, *elem)
	}

	return nil

}

func (r *mongoPostRepositoryStruct) Read(ctx context.Context, id int64, req *entity.Post) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	filter := bson.M{"id": id}

	err := collection.FindOne(ctx, filter).Decode(req)
	if err != nil {
		return err
	}

	return nil

}

func (r *mongoPostRepositoryStruct) Edit(ctx context.Context, id int64, req *entity.Post) error {

	current := entity.NewPost()

	collection := r.client.Database(r.db).Collection(r.collection)

	filter := bson.M{"id": id}

	err := collection.FindOne(ctx, filter).Decode(&current)
	if err != nil {
		return err
	}

	if req.Id != nil {
		current.Id = req.Id
	}

	if req.UserId != nil {
		current.UserId = req.UserId
	}

	if req.Title != nil {
		current.Title = req.Title
	}

	if req.Subtitle != nil {
		current.Subtitle = req.Subtitle
	}

	if req.Views != nil {
		current.Views = req.Views
	}

	if req.Tags != nil {
		current.Tags = req.Tags
	}

	if req.Created != nil {
		current.Created = req.Created
	}

	if req.Modified != nil {
		current.Modified = req.Modified
	}

	err = collection.FindOneAndReplace(ctx, filter, current).Decode(&current)
	if err != nil {
		return err
	}

	return nil

}

func (r *mongoPostRepositoryStruct) Add(ctx context.Context, req *entity.Post) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	_, err := collection.InsertOne(ctx, req)
	if err != nil {
		return err
	}

	return nil

}

func (r *mongoPostRepositoryStruct) Delete(ctx context.Context, id int64, req *entity.Post) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	filter := bson.M{"id": id}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil

}
