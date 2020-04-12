package entity

import bson "go.mongodb.org/mongo-driver/bson"

type bsonLucky struct {
	Thing  *string `bson:"thing,omitempty"`
	Thing2 *int64  `bson:"thing_2,omitempty"`
}

func (m *luckyStruct) marshalBSON() *bsonLucky {

	bsonStruct := bsonLucky{}

	bsonStruct.Thing = m.Thing()
	bsonStruct.Thing2 = m.Thing2()

	return &bsonStruct
}

func (m *luckyStruct) unmarshalBSON(bsonStruct *bsonLucky) {
	m.SetThing(bsonStruct.Thing)
	m.SetThing2(bsonStruct.Thing2)
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
