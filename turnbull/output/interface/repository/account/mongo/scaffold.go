package repository

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
	bson "go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

type mongoAccountRepositoryStruct struct {
	client     *mongo.Client
	db         string
	collection string
}

func (r *mongoAccountRepositoryStruct) Browse(ctx context.Context, req entity.Accounts) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return err
		}
	}

	for cursor.Next(ctx) {
		elem := entity.NewAccount()
		err := cursor.Decode(elem)
		if err != nil {
			return err
		}
		req.Append(elem)
	}

	return nil

}

func (r *mongoAccountRepositoryStruct) Read(ctx context.Context, id int64, req entity.Account) error {

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

func (r *mongoAccountRepositoryStruct) Edit(ctx context.Context, id int64, req entity.Account) error {

	current := entity.NewAccount()

	collection := r.client.Database(r.db).Collection(r.collection)

	filter := bson.M{"id": id}

	err := collection.FindOne(ctx, filter).Decode(&current)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return err
		}
	}

	if req.Id() != nil {
		current.SetId(req.Id())
	}

	if req.Name() != nil {
		current.SetName(req.Name())
	}

	if req.Created() != nil {
		current.SetCreated(req.Created())
	}

	if req.Modified() != nil {
		current.SetModified(req.Modified())
	}

	err = collection.FindOneAndReplace(ctx, filter, current).Decode(&current)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return err
		}
	}

	return nil

}

func (r *mongoAccountRepositoryStruct) Add(ctx context.Context, req entity.Account) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	_, err := collection.InsertOne(ctx, req)
	if err != nil {
		return err
	}

	return nil

}

func (r *mongoAccountRepositoryStruct) Delete(ctx context.Context, id int64, req entity.Account) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	filter := bson.M{"id": id}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil

}
