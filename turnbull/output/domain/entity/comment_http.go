package entity

import (
	"encoding/json"
	"time"
)

func (m *commentStruct) MarshalJSON() ([]byte, error) {
	type jsonStructPrivate struct {
		Id       *int64     `json:"id,omitempty"`
		PostId   *int64     `json:"post_id,omitempty"`
		UserId   *int64     `json:"user_id,omitempty"`
		Title    *string    `json:"title,omitempty"`
		Body     *string    `json:"body,omitempty"`
		Created  *time.Time `json:"created,omitempty"`
		Modified *time.Time `json:"modified,omitempty"`
	}
	jsonStruct := jsonStructPrivate{
		Body:     m.Body(),
		Created:  m.Created(),
		Id:       m.Id(),
		Modified: m.Modified(),
		PostId:   m.PostId(),
		Title:    m.Title(),
		UserId:   m.UserId(),
	}
	return json.Marshal(&jsonStruct)
}

func (m *commentStruct) UnmarshalJSON(data []byte) error {
	type jsonStructPrivate struct {
		Id       *int64     `json:"id,omitempty"`
		PostId   *int64     `json:"post_id,omitempty"`
		UserId   *int64     `json:"user_id,omitempty"`
		Title    *string    `json:"title,omitempty"`
		Body     *string    `json:"body,omitempty"`
		Created  *time.Time `json:"created,omitempty"`
		Modified *time.Time `json:"modified,omitempty"`
	}
	jsonStruct := jsonStructPrivate{}
	err := json.Unmarshal(data, &jsonStruct)
	if err != nil {
		return err
	}
	m.SetId(jsonStruct.Id)
	m.SetPostId(jsonStruct.PostId)
	m.SetUserId(jsonStruct.UserId)
	m.SetTitle(jsonStruct.Title)
	m.SetBody(jsonStruct.Body)
	m.SetCreated(jsonStruct.Created)
	m.SetModified(jsonStruct.Modified)
	return nil
}
