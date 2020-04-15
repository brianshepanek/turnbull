package entity

import bson "go.mongodb.org/mongo-driver/bson"

type bsonChannel struct {
	Model     *bsonModel `bson:"inline"`
	AccountId *int64     `bson:"account_id,omitempty"`
	Name      *string    `bson:"name,omitempty"`
}

func (m *channelStruct) marshalBSON() *bsonChannel {

	bsonStruct := bsonChannel{}

	bsonStruct.Model = m.model.marshalBSON()
	bsonStruct.AccountId = m.AccountId()
	bsonStruct.Name = m.Name()

	return &bsonStruct
}

func (m *channelStruct) unmarshalBSON(bsonStruct *bsonChannel) {
	m.model.unmarshalBSON(bsonStruct.Model)
	m.SetAccountId(bsonStruct.AccountId)
	m.SetName(bsonStruct.Name)
}

func (m *channelStruct) MarshalBSON() ([]byte, error) {
	return bson.Marshal(m.marshalBSON())
}

func (m *channelStruct) UnmarshalBSON(data []byte) error {

	bsonStruct := bsonChannel{}
	bsonStruct.Model = &bsonModel{}

	err := bson.Unmarshal(data, &bsonStruct)
	if err != nil {
		return err
	}

	m.unmarshalBSON(&bsonStruct)

	return nil

}
