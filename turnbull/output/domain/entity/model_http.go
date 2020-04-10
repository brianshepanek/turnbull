package entity

import (
	"encoding/json"
	"time"
)

type jsonModel struct {
	Id       *int64     `json:"id,omitempty"`
	Created  *time.Time `json:"created,omitempty"`
	Modified *time.Time `json:"modified,omitempty"`
}

func (m *modelStruct) marshalJSON() *jsonModel {

	jsonStruct := jsonModel{}

	jsonStruct.Id = m.Id()
	jsonStruct.Created = m.Created()
	jsonStruct.Modified = m.Modified()

	return &jsonStruct
}

func (m *modelStruct) unmarshalJSON(jsonStruct *jsonModel) {
	m.SetId(jsonStruct.Id)
	m.SetCreated(jsonStruct.Created)
	m.SetModified(jsonStruct.Modified)
}

func (m *modelStruct) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.marshalJSON())
}

func (m *modelStruct) UnmarshalJSON(data []byte) error {

	jsonStruct := jsonModel{}

	err := json.Unmarshal(data, &jsonStruct)
	if err != nil {
		return err
	}

	m.unmarshalJSON(&jsonStruct)

	return nil

}
