package entity

import (
	"context"
	"time"
)

type postStruct struct {
	id       *int64
	title    *string
	subtitle *string
	views    *int
	tags     *[]string
	created  *time.Time
	modified *time.Time
}

type postsStruct []postInterface

type postInterface interface {
	Id() *int64
	Title() *string
	Subtitle() *string
	Views() *int
	Tags() *[]string
	Created() *time.Time
	Modified() *time.Time
	SetId(id *int64)
	SetTitle(title *string)
	SetSubtitle(subtitle *string)
	SetViews(views *int)
	SetTags(tags *[]string)
	SetCreated(created *time.Time)
	SetModified(modified *time.Time)
	BeforeRead(ctx context.Context) error
	BeforeAdd(ctx context.Context) error
	SetAll(req postInterface)
	ToPrimary(ctx context.Context, req interface{}) (int64, error)
}
type postsInterface interface {
	Len() int
	Append(req postInterface)
	Elements() []postInterface
}

func (m *postsStruct) Len() int {
	if m != nil {
		return len(*m)
	}
	return 0
}
func (m *postsStruct) Append(req postInterface) {
	if m != nil {
		*m = append(*m, req)
	}
}
func (m *postsStruct) Elements() []postInterface {
	return *m
}
func (m *postStruct) Id() *int64 {
	return m.id
}

func (m *postStruct) Title() *string {
	return m.title
}

func (m *postStruct) Subtitle() *string {
	return m.subtitle
}

func (m *postStruct) Views() *int {
	return m.views
}

func (m *postStruct) Tags() *[]string {
	return m.tags
}

func (m *postStruct) Created() *time.Time {
	return m.created
}

func (m *postStruct) Modified() *time.Time {
	return m.modified
}

func (m *postStruct) SetId(id *int64) {
	m.id = id
}

func (m *postStruct) SetTitle(title *string) {
	m.title = title
}

func (m *postStruct) SetSubtitle(subtitle *string) {
	m.subtitle = subtitle
}

func (m *postStruct) SetViews(views *int) {
	m.views = views
}

func (m *postStruct) SetTags(tags *[]string) {
	m.tags = tags
}

func (m *postStruct) SetCreated(created *time.Time) {
	m.created = created
}

func (m *postStruct) SetModified(modified *time.Time) {
	m.modified = modified
}

func (m *postStruct) BeforeRead(ctx context.Context) error {
	return nil
}

func (m *postStruct) BeforeAdd(ctx context.Context) error {
	return nil
}

func (m *postStruct) SetAll(req postInterface) {
	m.SetId(req.Id())
	m.SetTitle(req.Title())
	m.SetSubtitle(req.Subtitle())
	m.SetViews(req.Views())
	m.SetTags(req.Tags())
	m.SetCreated(req.Created())
	m.SetModified(req.Modified())
}

func (m *postStruct) ToPrimary(ctx context.Context, req interface{}) (int64, error) {
	var resp int64
	return resp, nil
}
