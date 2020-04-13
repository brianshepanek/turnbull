package entity

import bson "go.mongodb.org/mongo-driver/bson"

type bsonUser struct {
	Model     *bsonModel `bson:"inline"`
	AccountId *int64     `bson:"account_id,omitempty"`
	FirstName *string    `bson:"first_name,omitempty"`
	LastName  *string    `bson:"last_name,omitempty"`
	Email     *string    `bson:"email,omitempty"`
}

func (m *userStruct) marshalBSON() *bsonUser {

	bsonStruct := bsonUser{}

	bsonStruct.Model = m.model.marshalBSON()
	bsonStruct.AccountId = m.AccountId()
	bsonStruct.FirstName = m.FirstName()
	bsonStruct.LastName = m.LastName()
	bsonStruct.Email = m.Email()

	return &bsonStruct
}

func (m *userStruct) unmarshalBSON(bsonStruct *bsonUser) {
	m.model.unmarshalBSON(bsonStruct.Model)
	m.SetAccountId(bsonStruct.AccountId)
	m.SetFirstName(bsonStruct.FirstName)
	m.SetLastName(bsonStruct.LastName)
	m.SetEmail(bsonStruct.Email)
}

func (m *userStruct) MarshalBSON() ([]byte, error) {
	return bson.Marshal(m.marshalBSON())
}

func (m *userStruct) UnmarshalBSON(data []byte) error {

	bsonStruct := bsonUser{}
	bsonStruct.Model = &bsonModel{}

	err := bson.Unmarshal(data, &bsonStruct)
	if err != nil {
		return err
	}

	m.unmarshalBSON(&bsonStruct)

	return nil

}
