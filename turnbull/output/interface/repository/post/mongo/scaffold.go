package repository

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
	bson "go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mongoPostRepositoryStruct struct {
	client     *mongo.Client
	db         string
	collection string
}
type post struct {
	entity.Post
}

func (m *post) MarshalBSON() ([]byte, error) {
	type bsonStructPrivate struct {
		Id       *int64     `bson:"id"`
		Title    *string    `bson:"title"`
		Subtitle *string    `bson:"subtitle"`
		Views    *int       `bson:"views"`
		Tags     *[]string  `bson:"tags"`
		Created  *time.Time `bson:"created"`
		Modified *time.Time `bson:"modified"`
	}
	bsonStruct := bsonStructPrivate{
		Created:  m.Created(),
		Id:       m.Id(),
		Modified: m.Modified(),
		Subtitle: m.Subtitle(),
		Tags:     m.Tags(),
		Title:    m.Title(),
		Views:    m.Views(),
	}
	return bson.Marshal(&bsonStruct)
}

func (m *post) UnmarshalBSON(data []byte) error {
	type bsonStructPrivate struct {
		Id       *int64     `bson:"id"`
		Title    *string    `bson:"title"`
		Subtitle *string    `bson:"subtitle"`
		Views    *int       `bson:"views"`
		Tags     *[]string  `bson:"tags"`
		Created  *time.Time `bson:"created"`
		Modified *time.Time `bson:"modified"`
	}
	bsonStruct := bsonStructPrivate{}
	err := bson.Unmarshal(data, &bsonStruct)
	if err != nil {
		return err
	}
	m.SetId(bsonStruct.Id)
	m.SetTitle(bsonStruct.Title)
	m.SetSubtitle(bsonStruct.Subtitle)
	m.SetViews(bsonStruct.Views)
	m.SetTags(bsonStruct.Tags)
	m.SetCreated(bsonStruct.Created)
	m.SetModified(bsonStruct.Modified)
	return nil
}

func (r *mongoPostRepositoryStruct) Browse(ctx context.Context, req entity.Posts) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return err
	}

	for cursor.Next(ctx) {
		elem := &post{entity.NewPost()}
		err := cursor.Decode(&elem)
		if err != nil {
			return err
		}
		req.Append(elem)
	}

	return nil

}

func (r *mongoPostRepositoryStruct) Read(ctx context.Context, id int64, req entity.Post) error {

	post := &post{req}

	collection := r.client.Database(r.db).Collection(r.collection)

	filter := bson.M{"id": id}

	err := collection.FindOne(ctx, filter).Decode(post)
	if err != nil {
		return err
	}

	return nil

}

func (r *mongoPostRepositoryStruct) Edit(ctx context.Context, id int64, req entity.Post) error {

	current := &post{entity.NewPost()}
	post := &post{req}

	collection := r.client.Database(r.db).Collection(r.collection)

	filter := bson.M{"id": id}

	err := collection.FindOne(ctx, filter).Decode(&current)
	if err != nil {
		return err
	}

	if post.Id() != nil {
		current.SetId(post.Id())
	}

	if post.Title() != nil {
		current.SetTitle(post.Title())
	}

	if post.Subtitle() != nil {
		current.SetSubtitle(post.Subtitle())
	}

	if post.Views() != nil {
		current.SetViews(post.Views())
	}

	if post.Tags() != nil {
		current.SetTags(post.Tags())
	}

	if post.Created() != nil {
		current.SetCreated(post.Created())
	}

	if post.Modified() != nil {
		current.SetModified(post.Modified())
	}

	err = collection.FindOneAndReplace(ctx, filter, current).Decode(&current)
	if err != nil {
		return err
	}

	return nil

}

func (r *mongoPostRepositoryStruct) Add(ctx context.Context, req entity.Post) error {

	post := &post{req}

	collection := r.client.Database(r.db).Collection(r.collection)

	_, err := collection.InsertOne(ctx, post)
	if err != nil {
		return err
	}

	return nil

}

func (r *mongoPostRepositoryStruct) Delete(ctx context.Context, id int64, req entity.Post) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	filter := bson.M{"id": id}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil

}
