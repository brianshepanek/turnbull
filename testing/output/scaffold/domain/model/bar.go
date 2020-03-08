package model

import (
	"encoding/json"
	"time"
)

type BarScaffoldStruct struct {
	slug     string
	title    string
	subtitle int64
	tags     []string
	created  time.Time
}

type BarsScaffoldStruct []BarScaffoldInterface

type BarScaffoldInterface interface {
	Slug() string
	Primary() string
	Title() string
	Subtitle() int64
	Tags() []string
	Created() time.Time
	SetSlug(slug string)
	SetTitle(title string)
	SetSubtitle(subtitle int64)
	SetTags(tags []string)
	SetCreated(created time.Time)
	SetAll(req BarScaffoldInterface)
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}
type BarsScaffoldInterface interface {
	Len() int
	Append(req BarScaffoldInterface)
	Elements() []BarScaffoldInterface
}

func NewBarScaffoldStruct() BarScaffoldInterface {
	return &BarScaffoldStruct{}
}

func NewBarsScaffoldStruct() BarsScaffoldInterface {
	return &BarsScaffoldStruct{}
}

func (m *BarsScaffoldStruct) Len() int {
	if m != nil {
		return len(*m)
	}
	return 0
}
func (m *BarsScaffoldStruct) Append(req BarScaffoldInterface) {
	if m != nil {
		*m = append(*m, req)
	}
}
func (m *BarsScaffoldStruct) Elements() []BarScaffoldInterface {
	return *m
}
func (m *BarScaffoldStruct) Slug() string {
	return m.slug
}

func (m *BarScaffoldStruct) Primary() string {
	return m.slug
}

func (m *BarScaffoldStruct) Title() string {
	return m.title
}

func (m *BarScaffoldStruct) Subtitle() int64 {
	return m.subtitle
}

func (m *BarScaffoldStruct) Tags() []string {
	return m.tags
}

func (m *BarScaffoldStruct) Created() time.Time {
	return m.created
}

func (m *BarScaffoldStruct) SetSlug(slug string) {
	m.slug = slug
}

func (m *BarScaffoldStruct) SetTitle(title string) {
	m.title = title
}

func (m *BarScaffoldStruct) SetSubtitle(subtitle int64) {
	m.subtitle = subtitle
}

func (m *BarScaffoldStruct) SetTags(tags []string) {
	m.tags = tags
}

func (m *BarScaffoldStruct) SetCreated(created time.Time) {
	m.created = created
}

func (m *BarScaffoldStruct) SetAll(req BarScaffoldInterface) {
	m.SetSlug(req.Slug())
	m.SetTitle(req.Title())
	m.SetSubtitle(req.Subtitle())
	m.SetTags(req.Tags())
	m.SetCreated(req.Created())
}

func (m *BarScaffoldStruct) MarshalJSON() ([]byte, error) {
	type jsonStructPrivate struct {
		Slug     string    `json:"slug"`
		Title    string    `json:"title"`
		Subtitle int64     `json:"subtitle"`
		Tags     []string  `json:"tags"`
		Created  time.Time `json:"created"`
	}
	jsonStruct := jsonStructPrivate{
		Created:  m.Created(),
		Slug:     m.Slug(),
		Subtitle: m.Subtitle(),
		Tags:     m.Tags(),
		Title:    m.Title(),
	}
	return json.Marshal(&jsonStruct)
}

func (m *BarScaffoldStruct) UnmarshalJSON(data []byte) error {
	type jsonStructPrivate struct {
		Slug     string    `json:"slug"`
		Title    string    `json:"title"`
		Subtitle int64     `json:"subtitle"`
		Tags     []string  `json:"tags"`
		Created  time.Time `json:"created"`
	}
	jsonStruct := jsonStructPrivate{}
	err := json.Unmarshal(data, &jsonStruct)
	if err != nil {
		return err
	}
	m.SetSlug(jsonStruct.Slug)
	m.SetTitle(jsonStruct.Title)
	m.SetSubtitle(jsonStruct.Subtitle)
	m.SetTags(jsonStruct.Tags)
	m.SetCreated(jsonStruct.Created)
	return nil
}
