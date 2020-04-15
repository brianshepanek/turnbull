package entity

import "encoding/json"

type jsonMessage struct {
	*jsonModel
	AccountId *int64  `json:"account_id,omitempty"`
	ChannelId *int64  `json:"channel_id,omitempty"`
	UserId    *int64  `json:"user_id,omitempty"`
	Message   *string `json:"message,omitempty"`
}

func (m *messageStruct) marshalJSON() *jsonMessage {

	jsonStruct := jsonMessage{}

	jsonStruct.jsonModel = m.model.marshalJSON()
	jsonStruct.AccountId = m.AccountId()
	jsonStruct.ChannelId = m.ChannelId()
	jsonStruct.UserId = m.UserId()
	jsonStruct.Message = m.Message()

	return &jsonStruct
}

func (m *messageStruct) unmarshalJSON(jsonStruct *jsonMessage) {
	m.model.unmarshalJSON(jsonStruct.jsonModel)
	m.SetAccountId(jsonStruct.AccountId)
	m.SetChannelId(jsonStruct.ChannelId)
	m.SetUserId(jsonStruct.UserId)
	m.SetMessage(jsonStruct.Message)
}

func (m *messageStruct) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.marshalJSON())
}

func (m *messageStruct) UnmarshalJSON(data []byte) error {

	jsonStruct := jsonMessage{}
	jsonStruct.jsonModel = &jsonModel{}

	err := json.Unmarshal(data, &jsonStruct)
	if err != nil {
		return err
	}

	m.unmarshalJSON(&jsonStruct)

	return nil

}
