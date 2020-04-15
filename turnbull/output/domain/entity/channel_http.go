package entity

import "encoding/json"

type jsonChannel struct {
	*jsonModel
	AccountId *int64  `json:"account_id,omitempty"`
	Name      *string `json:"name,omitempty"`
}

func (m *channelStruct) marshalJSON() *jsonChannel {

	jsonStruct := jsonChannel{}

	jsonStruct.jsonModel = m.model.marshalJSON()
	jsonStruct.AccountId = m.AccountId()
	jsonStruct.Name = m.Name()

	return &jsonStruct
}

func (m *channelStruct) unmarshalJSON(jsonStruct *jsonChannel) {
	m.model.unmarshalJSON(jsonStruct.jsonModel)
	m.SetAccountId(jsonStruct.AccountId)
	m.SetName(jsonStruct.Name)
}

func (m *channelStruct) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.marshalJSON())
}

func (m *channelStruct) UnmarshalJSON(data []byte) error {

	jsonStruct := jsonChannel{}
	jsonStruct.jsonModel = &jsonModel{}

	err := json.Unmarshal(data, &jsonStruct)
	if err != nil {
		return err
	}

	m.unmarshalJSON(&jsonStruct)

	return nil

}
