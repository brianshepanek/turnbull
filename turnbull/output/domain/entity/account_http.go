package entity

import "encoding/json"

type jsonAccount struct {
	*jsonModel
	*jsonLucky
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

func (m *accountStruct) marshalJSON() *jsonAccount {

	jsonStruct := jsonAccount{}

	jsonStruct.jsonModel = m.model.marshalJSON()
	jsonStruct.jsonLucky = m.lucky.marshalJSON()
	jsonStruct.Name = m.Name()
	jsonStruct.Email = m.Email()

	return &jsonStruct
}

func (m *accountStruct) unmarshalJSON(jsonStruct *jsonAccount) {
	m.model.unmarshalJSON(jsonStruct.jsonModel)
	m.lucky.unmarshalJSON(jsonStruct.jsonLucky)
	m.SetName(jsonStruct.Name)
	m.SetEmail(jsonStruct.Email)
}

func (m *accountStruct) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.marshalJSON())
}

func (m *accountStruct) UnmarshalJSON(data []byte) error {

	jsonStruct := jsonAccount{}
	jsonStruct.jsonModel = &jsonModel{}
	jsonStruct.jsonLucky = &jsonLucky{}

	err := json.Unmarshal(data, &jsonStruct)
	if err != nil {
		return err
	}

	m.unmarshalJSON(&jsonStruct)

	return nil

}
