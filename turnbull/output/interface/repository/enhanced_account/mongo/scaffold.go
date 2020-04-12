package repository

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
	bson "go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

type mongoEnhancedAccountRepositoryStruct struct {
	client     *mongo.Client
	db         string
	collection string
}

func (r *mongoEnhancedAccountRepositoryStruct) Browse(ctx context.Context, req entity.EnhancedAccounts) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return err
		}
	}

	for cursor.Next(ctx) {
		elem := entity.NewEnhancedAccount()
		err := cursor.Decode(elem)
		if err != nil {
			return err
		}
		req.Append(elem)
	}

	return nil

}

func (r *mongoEnhancedAccountRepositoryStruct) Read(ctx context.Context, id int64, req entity.EnhancedAccount) error {

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

func (r *mongoEnhancedAccountRepositoryStruct) Edit(ctx context.Context, id int64, req entity.EnhancedAccount) error {

	current := entity.NewEnhancedAccount()

	collection := r.client.Database(r.db).Collection(r.collection)

	filter := bson.M{"id": id}

	err := collection.FindOne(ctx, filter).Decode(current)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return err
		}
	}

	if req.Enhancement() != nil {
		current.SetEnhancement(req.Enhancement())
	}

	if req.Name() != nil {
		current.SetName(req.Name())
	}

	if req.Email() != nil {
		current.SetEmail(req.Email())
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

func (r *mongoEnhancedAccountRepositoryStruct) Add(ctx context.Context, req entity.EnhancedAccount) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	_, err := collection.InsertOne(ctx, req)
	if err != nil {
		return err
	}

	return nil

}

func (r *mongoEnhancedAccountRepositoryStruct) Delete(ctx context.Context, id int64, req entity.EnhancedAccount) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	filter := bson.M{"id": id}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil

}
