package repository

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
	bson "go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mongoFooRepositoryStruct struct {
	client     *mongo.Client
	db         string
	collection string
}
type foo struct {
	entity.Foo
}

func (m *foo) MarshalBSON() ([]byte, error) {
	type bsonStructPrivate struct {
		Id       string    `bson:"id"`
		Title    string    `bson:"title"`
		Subtitle string    `bson:"subtitle"`
		Int      int       `bson:"int"`
		Tags     []string  `bson:"tags"`
		Created  time.Time `bson:"created"`
	}
	bsonStruct := bsonStructPrivate{
		Created:  m.Created(),
		Id:       m.Id(),
		Int:      m.Int(),
		Subtitle: m.Subtitle(),
		Tags:     m.Tags(),
		Title:    m.Title(),
	}
	return bson.Marshal(&bsonStruct)
}

func (m *foo) UnmarshalBSON(data []byte) error {
	type bsonStructPrivate struct {
		Id       string    `bson:"id"`
		Title    string    `bson:"title"`
		Subtitle string    `bson:"subtitle"`
		Int      int       `bson:"int"`
		Tags     []string  `bson:"tags"`
		Created  time.Time `bson:"created"`
	}
	bsonStruct := bsonStructPrivate{}
	err := bson.Unmarshal(data, &bsonStruct)
	if err != nil {
		return err
	}
	m.SetId(bsonStruct.Id)
	m.SetTitle(bsonStruct.Title)
	m.SetSubtitle(bsonStruct.Subtitle)
	m.SetInt(bsonStruct.Int)
	m.SetTags(bsonStruct.Tags)
	m.SetCreated(bsonStruct.Created)
	return nil
}

func (r *mongoFooRepositoryStruct) Browse(ctx context.Context, query interface{}, req entity.Foos) error {
	collection := r.client.Database(r.db).Collection(r.collection)
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return err
	}
	for cursor.Next(ctx) {
		elem := &foo{entity.NewFoo()}
		err := cursor.Decode(&elem)
		if err != nil {
			return err
		}
		req.Append(elem)
	}
	return nil
}
func (r *mongoFooRepositoryStruct) Read(ctx context.Context, query interface{}, req entity.Foo) error {
	return nil
}
func (r *mongoFooRepositoryStruct) Edit(ctx context.Context, req entity.Foo) error {
	return nil
}
func (r *mongoFooRepositoryStruct) Add(ctx context.Context, req entity.Foo) error {
	foo := &foo{req}
	collection := r.client.Database(r.db).Collection(r.collection)
	res, err := collection.InsertOne(ctx, foo)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": res.InsertedID}
	err = collection.FindOne(ctx, filter).Decode(foo)
	return nil
}
func (r *mongoFooRepositoryStruct) Delete(ctx context.Context, req entity.Foo) error {
	return nil
}
