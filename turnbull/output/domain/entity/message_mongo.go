package entity

import bson "go.mongodb.org/mongo-driver/bson"

type bsonMessage struct {
	Model     *bsonModel `bson:"inline"`
	AccountId *int64     `bson:"account_id,omitempty"`
	ChannelId *int64     `bson:"channel_id,omitempty"`
	UserId    *int64     `bson:"user_id,omitempty"`
	Message   *string    `bson:"message,omitempty"`
}

func (m *messageStruct) marshalBSON() *bsonMessage {

	bsonStruct := bsonMessage{}

	bsonStruct.Model = m.model.marshalBSON()
	bsonStruct.AccountId = m.AccountId()
	bsonStruct.ChannelId = m.ChannelId()
	bsonStruct.UserId = m.UserId()
	bsonStruct.Message = m.Message()

	return &bsonStruct
}

func (m *messageStruct) unmarshalBSON(bsonStruct *bsonMessage) {
	m.model.unmarshalBSON(bsonStruct.Model)
	m.SetAccountId(bsonStruct.AccountId)
	m.SetChannelId(bsonStruct.ChannelId)
	m.SetUserId(bsonStruct.UserId)
	m.SetMessage(bsonStruct.Message)
}

func (m *messageStruct) MarshalBSON() ([]byte, error) {
	return bson.Marshal(m.marshalBSON())
}

func (m *messageStruct) UnmarshalBSON(data []byte) error {

	bsonStruct := bsonMessage{}
	bsonStruct.Model = &bsonModel{}

	err := bson.Unmarshal(data, &bsonStruct)
	if err != nil {
		return err
	}

	m.unmarshalBSON(&bsonStruct)

	return nil

}
