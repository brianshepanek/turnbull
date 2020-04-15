package repository

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
	bson "go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

type mongoChannelRepositoryStruct struct {
	client     *mongo.Client
	db         string
	collection string
}

func (r *mongoChannelRepositoryStruct) Browse(ctx context.Context, req entity.Channels) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return err
		}
	}

	for cursor.Next(ctx) {
		elem := entity.NewChannel()
		err := cursor.Decode(elem)
		if err != nil {
			return err
		}
		req.Append(elem)
	}

	return nil

}

func (r *mongoChannelRepositoryStruct) BrowseByAccountId(ctx context.Context, account_id int64, req entity.Channels) error {

	return nil

}

func (r *mongoChannelRepositoryStruct) Read(ctx context.Context, id int64, req entity.Channel) error {

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

func (r *mongoChannelRepositoryStruct) ReadByAccountIdAndName(ctx context.Context, account_id int64, name string, req entity.Channel) error {

	return nil

}

func (r *mongoChannelRepositoryStruct) Edit(ctx context.Context, id int64, req entity.Channel) error {

	current := entity.NewChannel()

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

	if req.Name() != nil {
		current.SetName(req.Name())
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

func (r *mongoChannelRepositoryStruct) Add(ctx context.Context, req entity.Channel) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	_, err := collection.InsertOne(ctx, req)
	if err != nil {
		return err
	}

	return nil

}

func (r *mongoChannelRepositoryStruct) Delete(ctx context.Context, id int64, req entity.Channel) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	filter := bson.M{"id": id}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil

}
