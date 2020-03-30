package entity

import (
	"context"
	"time"
)

type fooStruct struct {
	id       string
	title    string
	subtitle string
	int      int
	tags     []string
	created  time.Time
}

type foosStruct []fooInterface

type fooInterface interface {
	Id() string
	Title() string
	Subtitle() string
	Int() int
	Tags() []string
	Created() time.Time
	SetId(id string)
	SetTitle(title string)
	SetSubtitle(subtitle string)
	SetInt(int int)
	SetTags(tags []string)
	SetCreated(created time.Time)
	BeforeRead(ctx context.Context) error
	BeforeAdd(ctx context.Context) error
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

func (m *fooStruct) Title() string {
	return m.title
}

func (m *fooStruct) Subtitle() string {
	return m.subtitle
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

func (m *fooStruct) SetTitle(title string) {
	m.title = title
}

func (m *fooStruct) SetSubtitle(subtitle string) {
	m.subtitle = subtitle
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

func (m *fooStruct) BeforeRead(ctx context.Context) error {
	return nil
}

func (m *fooStruct) BeforeAdd(ctx context.Context) error {
	return nil
}
