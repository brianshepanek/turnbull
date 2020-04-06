package entity

import (
	bson "go.mongodb.org/mongo-driver/bson"
	"time"
)

func (m *postStruct) MarshalBSON() ([]byte, error) {
	type bsonStructPrivate struct {
		Id       *int64     `bson:"id"`
		UserId   *int64     `bson:"user_id"`
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
		UserId:   m.UserId,
		Views:    m.Views,
	}
	return bson.Marshal(&bsonStruct)
}

func (m *postStruct) UnmarshalBSON(data []byte) error {
	type bsonStructPrivate struct {
		Id       *int64     `bson:"id"`
		UserId   *int64     `bson:"user_id"`
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
	m.UserId = bsonStruct.UserId
	m.Title = bsonStruct.Title
	m.Subtitle = bsonStruct.Subtitle
	m.Views = bsonStruct.Views
	m.Tags = bsonStruct.Tags
	m.Created = bsonStruct.Created
	m.Modified = bsonStruct.Modified
	return nil
}
