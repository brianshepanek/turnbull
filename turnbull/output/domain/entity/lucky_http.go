package entity

import "encoding/json"

type jsonLucky struct {
	Thing *string `json:"thing,omitempty"`
}

func (m *luckyStruct) marshalJSON() *jsonLucky {

	jsonStruct := jsonLucky{}

	jsonStruct.Thing = m.Thing()

	return &jsonStruct
}

func (m *luckyStruct) unmarshalJSON(jsonStruct *jsonLucky) {
	m.SetThing(jsonStruct.Thing)
}

func (m *luckyStruct) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.marshalJSON())
}

func (m *luckyStruct) UnmarshalJSON(data []byte) error {

	jsonStruct := jsonLucky{}

	err := json.Unmarshal(data, &jsonStruct)
	if err != nil {
		return err
	}

	m.unmarshalJSON(&jsonStruct)

	return nil

}
