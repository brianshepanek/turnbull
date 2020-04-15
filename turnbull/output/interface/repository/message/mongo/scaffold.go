package repository

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
	bson "go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

type mongoMessageRepositoryStruct struct {
	client     *mongo.Client
	db         string
	collection string
}

func (r *mongoMessageRepositoryStruct) Browse(ctx context.Context, req entity.Messages) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return err
		}
	}

	for cursor.Next(ctx) {
		elem := entity.NewMessage()
		err := cursor.Decode(elem)
		if err != nil {
			return err
		}
		req.Append(elem)
	}

	return nil

}

func (r *mongoMessageRepositoryStruct) BrowseByAccountIdChannelId(ctx context.Context, account_id int64, channel_id int64, req entity.Messages) error {

	return nil

}

func (r *mongoMessageRepositoryStruct) Read(ctx context.Context, id int64, req entity.Message) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	filter := bson.M{"id": id}

	err := collection.FindOne(ctx, filter).Decode(req)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return err
		}
	}

	return nil

}

func (r *mongoMessageRepositoryStruct) Edit(ctx context.Context, id int64, req entity.Message) error {

	current := entity.NewMessage()

	collection := r.client.Database(r.db).Collection(r.collection)

	filter := bson.M{"id": id}

	err := collection.FindOne(ctx, filter).Decode(current)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return err
		}
	}

	if req.AccountId() != nil {
		current.SetAccountId(req.AccountId())
	}

	if req.ChannelId() != nil {
		current.SetChannelId(req.ChannelId())
	}

	if req.UserId() != nil {
		current.SetUserId(req.UserId())
	}

	if req.Message() != nil {
		current.SetMessage(req.Message())
	}

	if req.Id() != nil {
		current.SetId(req.Id())
	}

	if req.Created() != nil {
		current.SetCreated(req.Created())
	}

	if req.Modified() != nil {
		current.SetModified(req.Modified())
	}

	err = collection.FindOneAndReplace(ctx, filter, current).Decode(current)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return err
		}
	}

	return nil

}

func (r *mongoMessageRepositoryStruct) Add(ctx context.Context, req entity.Message) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	_, err := collection.InsertOne(ctx, req)
	if err != nil {
		return err
	}

	return nil

}

func (r *mongoMessageRepositoryStruct) Delete(ctx context.Context, id int64, req entity.Message) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	filter := bson.M{"id": id}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil

}
