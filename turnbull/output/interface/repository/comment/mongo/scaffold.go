package repository

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
	bson "go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

type mongoCommentRepositoryStruct struct {
	client     *mongo.Client
	db         string
	collection string
}

func (r *mongoCommentRepositoryStruct) Browse(ctx context.Context, req entity.Comments) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return err
	}

	for cursor.Next(ctx) {
		elem := entity.NewComment()
		err := cursor.Decode(elem)
		if err != nil {
			return err
		}
		req.Append(elem)
	}

	return nil

}

func (r *mongoCommentRepositoryStruct) Read(ctx context.Context, id int64, req entity.Comment) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	filter := bson.M{"id": id}

	err := collection.FindOne(ctx, filter).Decode(req)
	if err != nil {
		return err
	}

	return nil

}

func (r *mongoCommentRepositoryStruct) Edit(ctx context.Context, id int64, req entity.Comment) error {

	current := entity.NewComment()

	collection := r.client.Database(r.db).Collection(r.collection)

	filter := bson.M{"id": id}

	err := collection.FindOne(ctx, filter).Decode(&current)
	if err != nil {
		return err
	}

	if req.Id() != nil {
		current.SetId(req.Id())
	}

	if req.PostId() != nil {
		current.SetPostId(req.PostId())
	}

	if req.UserId() != nil {
		current.SetUserId(req.UserId())
	}

	if req.Title() != nil {
		current.SetTitle(req.Title())
	}

	if req.Body() != nil {
		current.SetBody(req.Body())
	}

	if req.Created() != nil {
		current.SetCreated(req.Created())
	}

	if req.Modified() != nil {
		current.SetModified(req.Modified())
	}

	err = collection.FindOneAndReplace(ctx, filter, current).Decode(&current)
	if err != nil {
		return err
	}

	return nil

}

func (r *mongoCommentRepositoryStruct) Add(ctx context.Context, req entity.Comment) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	_, err := collection.InsertOne(ctx, req)
	if err != nil {
		return err
	}

	return nil

}

func (r *mongoCommentRepositoryStruct) Delete(ctx context.Context, id int64, req entity.Comment) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	filter := bson.M{"id": id}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil

}
