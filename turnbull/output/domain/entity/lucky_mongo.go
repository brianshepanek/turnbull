package entity

import bson "go.mongodb.org/mongo-driver/bson"

type bsonLucky struct {
	Thing *string `bson:"thing,omitempty"`
}

func (m *luckyStruct) marshalBSON() *bsonLucky {

	bsonStruct := bsonLucky{}

	bsonStruct.Thing = m.Thing()

	return &bsonStruct
}

func (m *luckyStruct) unmarshalBSON(bsonStruct *bsonLucky) {
	m.SetThing(bsonStruct.Thing)
}

func (m *luckyStruct) MarshalBSON() ([]byte, error) {
	return bson.Marshal(m.marshalBSON())
}

func (m *luckyStruct) UnmarshalBSON(data []byte) error {

	bsonStruct := bsonLucky{}

	err := bson.Unmarshal(data, &bsonStruct)
	if err != nil {
		return err
	}

	m.unmarshalBSON(&bsonStruct)

	return nil

}
