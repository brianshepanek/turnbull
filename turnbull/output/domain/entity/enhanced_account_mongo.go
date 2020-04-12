package entity

import bson "go.mongodb.org/mongo-driver/bson"

type bsonEnhancedAccount struct {
	Account     *bsonAccount `bson:"inline"`
	Enhancement *string      `bson:"enhancement,omitempty"`
}

func (m *enhancedAccountStruct) marshalBSON() *bsonEnhancedAccount {

	bsonStruct := bsonEnhancedAccount{}

	bsonStruct.Account = m.account.marshalBSON()
	bsonStruct.Enhancement = m.Enhancement()

	return &bsonStruct
}

func (m *enhancedAccountStruct) unmarshalBSON(bsonStruct *bsonEnhancedAccount) {
	m.account.unmarshalBSON(bsonStruct.Account)
	m.SetEnhancement(bsonStruct.Enhancement)
}

func (m *enhancedAccountStruct) MarshalBSON() ([]byte, error) {
	return bson.Marshal(m.marshalBSON())
}

func (m *enhancedAccountStruct) UnmarshalBSON(data []byte) error {

	bsonStruct := bsonEnhancedAccount{}
	bsonStruct.Account = &bsonAccount{}

	err := bson.Unmarshal(data, &bsonStruct)
	if err != nil {
		return err
	}

	m.unmarshalBSON(&bsonStruct)

	return nil

}
