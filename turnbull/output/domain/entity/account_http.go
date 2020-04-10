package entity

import (
	"encoding/json"
	"time"
)

type jsonAccount struct {
	Id       *int64     `json:"id,omitempty"`
	Name     *string    `json:"name,omitempty"`
	Created  *time.Time `json:"created,omitempty"`
	Modified *time.Time `json:"modified,omitempty"`
}

func (m *accountStruct) marshalJSON() *jsonAccount {

	jsonStruct := jsonAccount{}

	jsonStruct.Id = m.Id()
	jsonStruct.Name = m.Name()
	jsonStruct.Created = m.Created()
	jsonStruct.Modified = m.Modified()

	return &jsonStruct
}

func (m *accountStruct) unmarshalJSON(jsonStruct jsonAccount) {
	m.SetId(jsonStruct.Id)
	m.SetName(jsonStruct.Name)
	m.SetCreated(jsonStruct.Created)
	m.SetModified(jsonStruct.Modified)
}

func (m *accountStruct) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.marshalJSON())
}

func (m *accountStruct) UnmarshalJSON(data []byte) error {

	jsonStruct := jsonAccount{}

	err := json.Unmarshal(data, &jsonStruct)
	if err != nil {
		return err
	}

	m.unmarshalJSON(jsonStruct)

	return nil

}
