package entity

import (
	bson "go.mongodb.org/mongo-driver/bson"
	"time"
)

type bsonModel struct {
	Id       *int64     `bson:"id,omitempty"`
	Created  *time.Time `bson:"created,omitempty"`
	Modified *time.Time `bson:"modified,omitempty"`
}

func (m *modelStruct) marshalBSON() *bsonModel {

	bsonStruct := bsonModel{}

	bsonStruct.Id = m.Id()
	bsonStruct.Created = m.Created()
	bsonStruct.Modified = m.Modified()

	return &bsonStruct
}

func (m *modelStruct) unmarshalBSON(bsonStruct *bsonModel) {
	m.SetId(bsonStruct.Id)
	m.SetCreated(bsonStruct.Created)
	m.SetModified(bsonStruct.Modified)
}

func (m *modelStruct) MarshalBSON() ([]byte, error) {
	return bson.Marshal(m.marshalBSON())
}

func (m *modelStruct) UnmarshalBSON(data []byte) error {

	bsonStruct := bsonModel{}

	err := bson.Unmarshal(data, &bsonStruct)
	if err != nil {
		return err
	}

	m.unmarshalBSON(&bsonStruct)

	return nil

}
