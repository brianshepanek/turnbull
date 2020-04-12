package entity

import bson "go.mongodb.org/mongo-driver/bson"

type bsonAccount struct {
	Model *bsonModel `bson:"inline"`
	Name  *string    `bson:"name,omitempty"`
	Email *string    `bson:"email,omitempty"`
}

func (m *accountStruct) marshalBSON() *bsonAccount {

	bsonStruct := bsonAccount{}

	bsonStruct.Model = m.model.marshalBSON()
	bsonStruct.Name = m.Name()
	bsonStruct.Email = m.Email()

	return &bsonStruct
}

func (m *accountStruct) unmarshalBSON(bsonStruct *bsonAccount) {
	m.model.unmarshalBSON(bsonStruct.Model)
	m.SetName(bsonStruct.Name)
	m.SetEmail(bsonStruct.Email)
}

func (m *accountStruct) MarshalBSON() ([]byte, error) {
	return bson.Marshal(m.marshalBSON())
}

func (m *accountStruct) UnmarshalBSON(data []byte) error {

	bsonStruct := bsonAccount{}
	bsonStruct.Model = &bsonModel{}

	err := bson.Unmarshal(data, &bsonStruct)
	if err != nil {
		return err
	}

	m.unmarshalBSON(&bsonStruct)

	return nil

}
