package repository

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
	bson "go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

type mongoUserRepositoryStruct struct {
	client     *mongo.Client
	db         string
	collection string
}

func (r *mongoUserRepositoryStruct) Browse(ctx context.Context, req entity.Users) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return err
		}
	}

	for cursor.Next(ctx) {
		elem := entity.NewUser()
		err := cursor.Decode(elem)
		if err != nil {
			return err
		}
		req.Append(elem)
	}

	return nil

}

func (r *mongoUserRepositoryStruct) BrowseByAccountId(ctx context.Context, id int64, req entity.Users) error {

	return nil

}

func (r *mongoUserRepositoryStruct) Read(ctx context.Context, id int64, req entity.User) error {

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

func (r *mongoUserRepositoryStruct) Edit(ctx context.Context, id int64, req entity.User) error {

	current := entity.NewUser()

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

	if req.FirstName() != nil {
		current.SetFirstName(req.FirstName())
	}

	if req.LastName() != nil {
		current.SetLastName(req.LastName())
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

func (r *mongoUserRepositoryStruct) Add(ctx context.Context, req entity.User) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	_, err := collection.InsertOne(ctx, req)
	if err != nil {
		return err
	}

	return nil

}

func (r *mongoUserRepositoryStruct) Delete(ctx context.Context, id int64, req entity.User) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	filter := bson.M{"id": id}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil

}
