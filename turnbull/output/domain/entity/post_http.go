package entity

import (
	"encoding/json"
	"time"
)

func (m *postStruct) MarshalJSON() ([]byte, error) {
	type jsonStructPrivate struct {
		Id       *int64     `json:"id,omitempty"`
		UserId   *int64     `json:"user_id,omitempty"`
		Title    *string    `json:"title,omitempty"`
		Subtitle *string    `json:"subtitle,omitempty"`
		Views    *int       `json:"views,omitempty"`
		Tags     *[]string  `json:"tags,omitempty"`
		Created  *time.Time `json:"created,omitempty"`
		Modified *time.Time `json:"modified,omitempty"`
	}
	jsonStruct := jsonStructPrivate{
		Created:  m.Created,
		Id:       m.Id,
		Modified: m.Modified,
		Subtitle: m.Subtitle,
		Tags:     m.Tags,
		Title:    m.Title,
		UserId:   m.UserId,
		Views:    m.Views,
	}
	return json.Marshal(&jsonStruct)
}

func (m *postStruct) UnmarshalJSON(data []byte) error {
	type jsonStructPrivate struct {
		Id       *int64     `json:"id,omitempty"`
		UserId   *int64     `json:"user_id,omitempty"`
		Title    *string    `json:"title,omitempty"`
		Subtitle *string    `json:"subtitle,omitempty"`
		Views    *int       `json:"views,omitempty"`
		Tags     *[]string  `json:"tags,omitempty"`
		Created  *time.Time `json:"created,omitempty"`
		Modified *time.Time `json:"modified,omitempty"`
	}
	jsonStruct := jsonStructPrivate{}
	err := json.Unmarshal(data, &jsonStruct)
	if err != nil {
		return err
	}
	m.Id = jsonStruct.Id
	m.UserId = jsonStruct.UserId
	m.Title = jsonStruct.Title
	m.Subtitle = jsonStruct.Subtitle
	m.Views = jsonStruct.Views
	m.Tags = jsonStruct.Tags
	m.Created = jsonStruct.Created
	m.Modified = jsonStruct.Modified
	return nil
}
