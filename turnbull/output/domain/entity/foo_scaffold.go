package entity

import (
	"encoding/json"
	"time"
)

type fooStruct struct {
	id      string
	string  string
	int     int
	tags    []string
	created time.Time
}

type foosStruct []fooInterface

type fooInterface interface {
	Id() string
	String() string
	Int() int
	Tags() []string
	Created() time.Time
	SetId(id string)
	SetString(string string)
	SetInt(int int)
	SetTags(tags []string)
	SetCreated(created time.Time)
	SetAll(req fooInterface)
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}
type foosInterface interface {
	Len() int
	Append(req fooInterface)
	Elements() []fooInterface
}

func (m *foosStruct) Len() int {
	if m != nil {
		return len(*m)
	}
	return 0
}
func (m *foosStruct) Append(req fooInterface) {
	if m != nil {
		*m = append(*m, req)
	}
}
func (m *foosStruct) Elements() []fooInterface {
	return *m
}
func (m *fooStruct) Id() string {
	return m.id
}

func (m *fooStruct) String() string {
	return m.string
}

func (m *fooStruct) Int() int {
	return m.int
}

func (m *fooStruct) Tags() []string {
	return m.tags
}

func (m *fooStruct) Created() time.Time {
	return m.created
}

func (m *fooStruct) SetId(id string) {
	m.id = id
}

func (m *fooStruct) SetString(string string) {
	m.string = string
}

func (m *fooStruct) SetInt(int int) {
	m.int = int
}

func (m *fooStruct) SetTags(tags []string) {
	m.tags = tags
}

func (m *fooStruct) SetCreated(created time.Time) {
	m.created = created
}

func (m *fooStruct) SetAll(req fooInterface) {
	m.SetId(req.Id())
	m.SetString(req.String())
	m.SetInt(req.Int())
	m.SetTags(req.Tags())
	m.SetCreated(req.Created())
}

func (m *fooStruct) MarshalJSON() ([]byte, error) {
	type jsonStructPrivate struct {
		Id      string    `json:"id"`
		String  string    `json:"string"`
		Int     int       `json:"int"`
		Tags    []string  `json:"tags"`
		Created time.Time `json:"created"`
	}
	jsonStruct := jsonStructPrivate{
		Created: m.Created(),
		Id:      m.Id(),
		Int:     m.Int(),
		String:  m.String(),
		Tags:    m.Tags(),
	}
	return json.Marshal(&jsonStruct)
}

func (m *fooStruct) UnmarshalJSON(data []byte) error {
	type jsonStructPrivate struct {
		Id      string    `json:"id"`
		String  string    `json:"string"`
		Int     int       `json:"int"`
		Tags    []string  `json:"tags"`
		Created time.Time `json:"created"`
	}
	jsonStruct := jsonStructPrivate{}
	err := json.Unmarshal(data, &jsonStruct)
	if err != nil {
		return err
	}
	m.SetId(jsonStruct.Id)
	m.SetString(jsonStruct.String)
	m.SetInt(jsonStruct.Int)
	m.SetTags(jsonStruct.Tags)
	m.SetCreated(jsonStruct.Created)
	return nil
}
