package entity

import "encoding/json"

type jsonUser struct {
	*jsonModel
	AccountId *int64  `json:"account_id,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Email     *string `json:"email,omitempty"`
}

func (m *userStruct) marshalJSON() *jsonUser {

	jsonStruct := jsonUser{}

	jsonStruct.jsonModel = m.model.marshalJSON()
	jsonStruct.AccountId = m.AccountId()
	jsonStruct.FirstName = m.FirstName()
	jsonStruct.LastName = m.LastName()
	jsonStruct.Email = m.Email()

	return &jsonStruct
}

func (m *userStruct) unmarshalJSON(jsonStruct *jsonUser) {
	m.model.unmarshalJSON(jsonStruct.jsonModel)
	m.SetAccountId(jsonStruct.AccountId)
	m.SetFirstName(jsonStruct.FirstName)
	m.SetLastName(jsonStruct.LastName)
	m.SetEmail(jsonStruct.Email)
}

func (m *userStruct) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.marshalJSON())
}

func (m *userStruct) UnmarshalJSON(data []byte) error {

	jsonStruct := jsonUser{}
	jsonStruct.jsonModel = &jsonModel{}

	err := json.Unmarshal(data, &jsonStruct)
	if err != nil {
		return err
	}

	m.unmarshalJSON(&jsonStruct)

	return nil

}
