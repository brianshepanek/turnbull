package entity

import (
	"context"
	"time"
)

type fooStruct struct {
	id       int64
	title    string
	subtitle string
	views    int
	tags     []string
	created  time.Time
	modified time.Time
}

type foosStruct []fooInterface

type fooInterface interface {
	Id() int64
	Title() string
	Subtitle() string
	Views() int
	Tags() []string
	Created() time.Time
	Modified() time.Time
	SetId(id int64)
	SetTitle(title string)
	SetSubtitle(subtitle string)
	SetViews(views int)
	SetTags(tags []string)
	SetCreated(created time.Time)
	SetModified(modified time.Time)
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
func (m *fooStruct) Id() int64 {
	return m.id
}

func (m *fooStruct) Title() string {
	return m.title
}

func (m *fooStruct) Subtitle() string {
	return m.subtitle
}

func (m *fooStruct) Views() int {
	return m.views
}

func (m *fooStruct) Tags() []string {
	return m.tags
}

func (m *fooStruct) Created() time.Time {
	return m.created
}

func (m *fooStruct) Modified() time.Time {
	return m.modified
}

func (m *fooStruct) SetId(id int64) {
	m.id = id
}

func (m *fooStruct) SetTitle(title string) {
	m.title = title
}

func (m *fooStruct) SetSubtitle(subtitle string) {
	m.subtitle = subtitle
}

func (m *fooStruct) SetViews(views int) {
	m.views = views
}

func (m *fooStruct) SetTags(tags []string) {
	m.tags = tags
}

func (m *fooStruct) SetCreated(created time.Time) {
	m.created = created
}

func (m *fooStruct) SetModified(modified time.Time) {
	m.modified = modified
}

func (m *fooStruct) BeforeRead(ctx context.Context) error {
	return nil
}

func (m *fooStruct) BeforeAdd(ctx context.Context) error {
	return nil
}
