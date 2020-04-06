package entity

import (
	bson "go.mongodb.org/mongo-driver/bson"
	"time"
)

func (m *commentStruct) MarshalBSON() ([]byte, error) {
	type bsonStructPrivate struct {
		Id       *int64     `bson:"id"`
		PostId   *int64     `bson:"post_id"`
		UserId   *int64     `bson:"user_id"`
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
		UserId:   m.UserId(),
	}
	return bson.Marshal(&bsonStruct)
}

func (m *commentStruct) UnmarshalBSON(data []byte) error {
	type bsonStructPrivate struct {
		Id       *int64     `bson:"id"`
		PostId   *int64     `bson:"post_id"`
		UserId   *int64     `bson:"user_id"`
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
	m.SetUserId(bsonStruct.UserId)
	m.SetTitle(bsonStruct.Title)
	m.SetBody(bsonStruct.Body)
	m.SetCreated(bsonStruct.Created)
	m.SetModified(bsonStruct.Modified)
	return nil
}
