package entity

import (
	"encoding/json"
	"time"
)

type FooScaffoldStruct struct {
	string  string
	int     int
	tags    []string
	created time.Time
}

type FoosScaffoldStruct []FooScaffoldInterface

type FooScaffoldInterface interface {
	String() string
	Int() int
	Tags() []string
	Created() time.Time
	SetString(string string)
	SetInt(int int)
	SetTags(tags []string)
	SetCreated(created time.Time)
	SetAll(req FooScaffoldInterface)
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}
type FoosScaffoldInterface interface {
	Len() int
	Append(req FooScaffoldInterface)
	Elements() []FooScaffoldInterface
}

func NewFooScaffoldStruct() FooScaffoldInterface {
	return &FooScaffoldStruct{}
}

func NewFoosScaffoldStruct() FoosScaffoldInterface {
	return &FoosScaffoldStruct{}
}

func (m *FoosScaffoldStruct) Len() int {
	if m != nil {
		return len(*m)
	}
	return 0
}
func (m *FoosScaffoldStruct) Append(req FooScaffoldInterface) {
	if m != nil {
		*m = append(*m, req)
	}
}
func (m *FoosScaffoldStruct) Elements() []FooScaffoldInterface {
	return *m
}
func (m *FooScaffoldStruct) String() string {
	return m.string
}

func (m *FooScaffoldStruct) Int() int {
	return m.int
}

func (m *FooScaffoldStruct) Tags() []string {
	return m.tags
}

func (m *FooScaffoldStruct) Created() time.Time {
	return m.created
}

func (m *FooScaffoldStruct) SetString(string string) {
	m.string = string
}

func (m *FooScaffoldStruct) SetInt(int int) {
	m.int = int
}

func (m *FooScaffoldStruct) SetTags(tags []string) {
	m.tags = tags
}

func (m *FooScaffoldStruct) SetCreated(created time.Time) {
	m.created = created
}

func (m *FooScaffoldStruct) SetAll(req FooScaffoldInterface) {
	m.SetString(req.String())
	m.SetInt(req.Int())
	m.SetTags(req.Tags())
	m.SetCreated(req.Created())
}

func (m *FooScaffoldStruct) MarshalJSON() ([]byte, error) {
	type jsonStructPrivate struct {
		String  string    `json:"string"`
		Int     int       `json:"int"`
		Tags    []string  `json:"tags"`
		Created time.Time `json:"created"`
	}
	jsonStruct := jsonStructPrivate{
		Created: m.Created(),
		Int:     m.Int(),
		String:  m.String(),
		Tags:    m.Tags(),
	}
	return json.Marshal(&jsonStruct)
}

func (m *FooScaffoldStruct) UnmarshalJSON(data []byte) error {
	type jsonStructPrivate struct {
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
	m.SetString(jsonStruct.String)
	m.SetInt(jsonStruct.Int)
	m.SetTags(jsonStruct.Tags)
	m.SetCreated(jsonStruct.Created)
	return nil
}
