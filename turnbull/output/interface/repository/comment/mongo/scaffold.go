package repository

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
	bson "go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mongoCommentRepositoryStruct struct {
	client     *mongo.Client
	db         string
	collection string
}
type comment struct {
	entity.Comment
}

func (m *comment) MarshalBSON() ([]byte, error) {
	type bsonStructPrivate struct {
		Id       *int64     `bson:"id"`
		PostId   *int64     `bson:"post_id"`
		Title    *string    `bson:"title"`
		Body     *string    `bson:"body"`
		Created  *time.Time `bson:"created"`
		Modified *time.Time `bson:"modified"`
	}
	bsonStruct := bsonStructPrivate{
		Body:     m.Body(),
		Created:  m.Created(),
		Id:       m.Id(),
		Modified: m.Modified(),
		PostId:   m.PostId(),
		Title:    m.Title(),
	}
	return bson.Marshal(&bsonStruct)
}

func (m *comment) UnmarshalBSON(data []byte) error {
	type bsonStructPrivate struct {
		Id       *int64     `bson:"id"`
		PostId   *int64     `bson:"post_id"`
		Title    *string    `bson:"title"`
		Body     *string    `bson:"body"`
		Created  *time.Time `bson:"created"`
		Modified *time.Time `bson:"modified"`
	}
	bsonStruct := bsonStructPrivate{}
	err := bson.Unmarshal(data, &bsonStruct)
	if err != nil {
		return err
	}
	m.SetId(bsonStruct.Id)
	m.SetPostId(bsonStruct.PostId)
	m.SetTitle(bsonStruct.Title)
	m.SetBody(bsonStruct.Body)
	m.SetCreated(bsonStruct.Created)
	m.SetModified(bsonStruct.Modified)
	return nil
}

func (r *mongoCommentRepositoryStruct) Browse(ctx context.Context, req entity.Comments) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return err
	}

	for cursor.Next(ctx) {
		elem := &comment{entity.NewComment()}
		err := cursor.Decode(&elem)
		if err != nil {
			return err
		}
		req.Append(elem)
	}

	return nil

}

func (r *mongoCommentRepositoryStruct) Read(ctx context.Context, id int64, req entity.Comment) error {

	comment := &comment{req}

	collection := r.client.Database(r.db).Collection(r.collection)

	filter := bson.M{"id": id}

	err := collection.FindOne(ctx, filter).Decode(comment)
	if err != nil {
		return err
	}

	return nil

}

func (r *mongoCommentRepositoryStruct) Edit(ctx context.Context, id int64, req entity.Comment) error {

	current := &comment{entity.NewComment()}
	comment := &comment{req}

	collection := r.client.Database(r.db).Collection(r.collection)

	filter := bson.M{"id": id}

	err := collection.FindOne(ctx, filter).Decode(&current)
	if err != nil {
		return err
	}

	if comment.Id() != nil {
		current.SetId(comment.Id())
	}

	if comment.PostId() != nil {
		current.SetPostId(comment.PostId())
	}

	if comment.Title() != nil {
		current.SetTitle(comment.Title())
	}

	if comment.Body() != nil {
		current.SetBody(comment.Body())
	}

	if comment.Created() != nil {
		current.SetCreated(comment.Created())
	}

	if comment.Modified() != nil {
		current.SetModified(comment.Modified())
	}

	err = collection.FindOneAndReplace(ctx, filter, current).Decode(&current)
	if err != nil {
		return err
	}

	return nil

}

func (r *mongoCommentRepositoryStruct) Add(ctx context.Context, req entity.Comment) error {

	comment := &comment{req}

	collection := r.client.Database(r.db).Collection(r.collection)

	_, err := collection.InsertOne(ctx, comment)
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
