package entity

import "encoding/json"

type jsonLucky struct {
	Thing  *string `json:"thing,omitempty"`
	Thing2 *int64  `json:"thing_2,omitempty"`
}

func (m *luckyStruct) marshalJSON() *jsonLucky {

	jsonStruct := jsonLucky{}

	jsonStruct.Thing = m.Thing()
	jsonStruct.Thing2 = m.Thing2()

	return &jsonStruct
}

func (m *luckyStruct) unmarshalJSON(jsonStruct *jsonLucky) {
	m.SetThing(jsonStruct.Thing)
	m.SetThing2(jsonStruct.Thing2)
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
