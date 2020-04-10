package entity

import (
	bson "go.mongodb.org/mongo-driver/bson"
	"time"
)

func (m *accountStruct) MarshalBSON() ([]byte, error) {
	type bsonStructPrivate struct {
		Id       *int64     `bson:"id"`
		Name     *string    `bson:"name"`
		Created  *time.Time `bson:"created"`
		Modified *time.Time `bson:"modified"`
	}
	bsonStruct := bsonStructPrivate{
		Created:  m.Created(),
		Id:       m.Id(),
		Modified: m.Modified(),
		Name:     m.Name(),
	}
	return bson.Marshal(&bsonStruct)
}

func (m *accountStruct) UnmarshalBSON(data []byte) error {
	type bsonStructPrivate struct {
		Id       *int64     `bson:"id"`
		Name     *string    `bson:"name"`
		Created  *time.Time `bson:"created"`
		Modified *time.Time `bson:"modified"`
	}
	bsonStruct := bsonStructPrivate{}
	err := bson.Unmarshal(data, &bsonStruct)
	if err != nil {
		return err
	}
	m.SetId(bsonStruct.Id)
	m.SetName(bsonStruct.Name)
	m.SetCreated(bsonStruct.Created)
	m.SetModified(bsonStruct.Modified)
	return nil
}
