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
type postLocal struct {
	*entity.Post
}

func (m *postLocal) MarshalBSON() ([]byte, error) {
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
		Created:  m.Created,
		Id:       m.Id,
		Modified: m.Modified,
		Subtitle: m.Subtitle,
		Tags:     m.Tags,
		Title:    m.Title,
		Views:    m.Views,
	}
	return bson.Marshal(&bsonStruct)
}

func (m *postLocal) UnmarshalBSON(data []byte) error {
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
	m.Id = bsonStruct.Id
	m.Title = bsonStruct.Title
	m.Subtitle = bsonStruct.Subtitle
	m.Views = bsonStruct.Views
	m.Tags = bsonStruct.Tags
	m.Created = bsonStruct.Created
	m.Modified = bsonStruct.Modified
	return nil
}

func (r *mongoPostRepositoryStruct) Browse(ctx context.Context, req *[]entity.Post) error {

	collection := r.client.Database(r.db).Collection(r.collection)

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return err
	}

	for cursor.Next(ctx) {
		elem := &postLocal{entity.NewPost()}
		err := cursor.Decode(&elem)
		if err != nil {
			return err
		}
		*req = append(*req, *elem.Post)
	}

	return nil

}

func (r *mongoPostRepositoryStruct) Read(ctx context.Context, id int64, req *entity.Post) error {

	postLocal := &postLocal{req}

	collection := r.client.Database(r.db).Collection(r.collection)

	filter := bson.M{"id": id}

	err := collection.FindOne(ctx, filter).Decode(postLocal)
	if err != nil {
		return err
	}

	return nil

}

func (r *mongoPostRepositoryStruct) Edit(ctx context.Context, id int64, req *entity.Post) error {

	current := &postLocal{entity.NewPost()}
	postLocal := &postLocal{req}

	collection := r.client.Database(r.db).Collection(r.collection)

	filter := bson.M{"id": id}

	err := collection.FindOne(ctx, filter).Decode(&current)
	if err != nil {
		return err
	}

	if postLocal.Id != nil {
		current.Id = postLocal.Id
	}

	if postLocal.Title != nil {
		current.Title = postLocal.Title
	}

	if postLocal.Subtitle != nil {
		current.Subtitle = postLocal.Subtitle
	}

	if postLocal.Views != nil {
		current.Views = postLocal.Views
	}

	if postLocal.Tags != nil {
		current.Tags = postLocal.Tags
	}

	if postLocal.Created != nil {
		current.Created = postLocal.Created
	}

	if postLocal.Modified != nil {
		current.Modified = postLocal.Modified
	}

	err = collection.FindOneAndReplace(ctx, filter, current).Decode(&current)
	if err != nil {
		return err
	}

	return nil

}

func (r *mongoPostRepositoryStruct) Add(ctx context.Context, req *entity.Post) error {

	postLocal := &postLocal{req}

	collection := r.client.Database(r.db).Collection(r.collection)

	_, err := collection.InsertOne(ctx, postLocal)
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
