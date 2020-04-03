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
type commentLocal struct {
	*entity.Comment
}

func (m *commentLocal) MarshalBSON() ([]byte, error) {
	type bsonStructPrivate struct {
		Id       *int64     `bson:"id"`
		PostId   *int64     `bson:"post_id"`
		Title    *string    `bson:"title"`
		Body     *string    `bson:"body"`
		Created  *time.Time `bson:"created"`
		Modified *time.Time `bson:"modified"`
	}
	bsonStruct := bsonStructPrivate{
		Body:     m.Body,
		Created:  m.Created,
		Id:       m.Id,
		Modified: m.Modified,
		PostId:   m.PostId,
		Title:    m.Title,
	}
	return bson.Marshal(&bsonStruct)
}

func (m *commentLocal) UnmarshalBSON(data []byte) error {
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
	m.Id = bsonStruct.Id
	m.PostId = bsonStruct.PostId
	m.Title = bsonStruct.Title
	m.Body = bsonStruct.Body
	m.Created = bsonStruct.Created
	m.Modified = bsonStruct.Modified
	return nil
}

func (r *mongoCommentRepositoryStruct) Browse(ctx context.Context, req *[]entity.Comment) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return err
	}

	for cursor.Next(ctx) {
		elem := &commentLocal{entity.NewComment()}
		err := cursor.Decode(&elem)
		if err != nil {
			return err
		}
		*req = append(*req, *elem.Comment)
	}

	return nil

}

func (r *mongoCommentRepositoryStruct) Read(ctx context.Context, id int64, req *entity.Comment) error {

	commentLocal := &commentLocal{req}

	collection := r.client.Database(r.db).Collection(r.collection)

	filter := bson.M{"id": id}

	err := collection.FindOne(ctx, filter).Decode(commentLocal)
	if err != nil {
		return err
	}

	return nil

}

func (r *mongoCommentRepositoryStruct) Edit(ctx context.Context, id int64, req *entity.Comment) error {

	current := &commentLocal{entity.NewComment()}
	commentLocal := &commentLocal{req}

	collection := r.client.Database(r.db).Collection(r.collection)

	filter := bson.M{"id": id}

	err := collection.FindOne(ctx, filter).Decode(&current)
	if err != nil {
		return err
	}

	if commentLocal.Id != nil {
		current.Id = commentLocal.Id
	}

	if commentLocal.PostId != nil {
		current.PostId = commentLocal.PostId
	}

	if commentLocal.Title != nil {
		current.Title = commentLocal.Title
	}

	if commentLocal.Body != nil {
		current.Body = commentLocal.Body
	}

	if commentLocal.Created != nil {
		current.Created = commentLocal.Created
	}

	if commentLocal.Modified != nil {
		current.Modified = commentLocal.Modified
	}

	err = collection.FindOneAndReplace(ctx, filter, current).Decode(&current)
	if err != nil {
		return err
	}

	return nil

}

func (r *mongoCommentRepositoryStruct) Add(ctx context.Context, req *entity.Comment) error {

	commentLocal := &commentLocal{req}

	collection := r.client.Database(r.db).Collection(r.collection)

	_, err := collection.InsertOne(ctx, commentLocal)
	if err != nil {
		return err
	}

	return nil

}

func (r *mongoCommentRepositoryStruct) Delete(ctx context.Context, id int64, req *entity.Comment) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	filter := bson.M{"id": id}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil

}
