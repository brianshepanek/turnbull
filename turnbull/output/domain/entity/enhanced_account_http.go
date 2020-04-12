package entity

import "encoding/json"

type jsonEnhancedAccount struct {
	*jsonAccount
	Enhancement *string `json:"enhancement,omitempty"`
}

func (m *enhancedAccountStruct) marshalJSON() *jsonEnhancedAccount {

	jsonStruct := jsonEnhancedAccount{}

	jsonStruct.jsonAccount = m.account.marshalJSON()
	jsonStruct.Enhancement = m.Enhancement()

	return &jsonStruct
}

func (m *enhancedAccountStruct) unmarshalJSON(jsonStruct *jsonEnhancedAccount) {
	m.account.unmarshalJSON(jsonStruct.jsonAccount)
	m.SetEnhancement(jsonStruct.Enhancement)
}

func (m *enhancedAccountStruct) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.marshalJSON())
}

func (m *enhancedAccountStruct) UnmarshalJSON(data []byte) error {

	jsonStruct := jsonEnhancedAccount{}
	jsonStruct.jsonAccount = &jsonAccount{}

	err := json.Unmarshal(data, &jsonStruct)
	if err != nil {
		return err
	}

	m.unmarshalJSON(&jsonStruct)

	return nil

}
