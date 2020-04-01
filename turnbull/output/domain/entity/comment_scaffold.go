package entity

import (
	"context"
	"time"
)

type commentStruct struct {
	id       *int64
	postId   *int64
	title    *string
	body     *string
	created  *time.Time
	modified *time.Time
}

type commentsStruct []commentInterface

type commentInterface interface {
	Id() *int64
	PostId() *int64
	Title() *string
	Body() *string
	Created() *time.Time
	Modified() *time.Time
	SetId(id *int64)
	SetPostId(postId *int64)
	SetTitle(title *string)
	SetBody(body *string)
	SetCreated(created *time.Time)
	SetModified(modified *time.Time)
	BeforeRead(ctx context.Context) error
	BeforeAdd(ctx context.Context) error
	SetAll(req commentInterface)
	ToPrimary(ctx context.Context, req interface{}) (int64, error)
}
type commentsInterface interface {
	Len() int
	Append(req commentInterface)
	Elements() []commentInterface
}

func (m *commentsStruct) Len() int {
	if m != nil {
		return len(*m)
	}
	return 0
}
func (m *commentsStruct) Append(req commentInterface) {
	if m != nil {
		*m = append(*m, req)
	}
}
func (m *commentsStruct) Elements() []commentInterface {
	return *m
}
func (m *commentStruct) Id() *int64 {
	return m.id
}

func (m *commentStruct) PostId() *int64 {
	return m.postId
}

func (m *commentStruct) Title() *string {
	return m.title
}

func (m *commentStruct) Body() *string {
	return m.body
}

func (m *commentStruct) Created() *time.Time {
	return m.created
}

func (m *commentStruct) Modified() *time.Time {
	return m.modified
}

func (m *commentStruct) SetId(id *int64) {
	m.id = id
}

func (m *commentStruct) SetPostId(postId *int64) {
	m.postId = postId
}

func (m *commentStruct) SetTitle(title *string) {
	m.title = title
}

func (m *commentStruct) SetBody(body *string) {
	m.body = body
}

func (m *commentStruct) SetCreated(created *time.Time) {
	m.created = created
}

func (m *commentStruct) SetModified(modified *time.Time) {
	m.modified = modified
}

func (m *commentStruct) BeforeRead(ctx context.Context) error {
	return nil
}

func (m *commentStruct) BeforeAdd(ctx context.Context) error {
	return nil
}

func (m *commentStruct) SetAll(req commentInterface) {
	m.SetId(req.Id())
	m.SetPostId(req.PostId())
	m.SetTitle(req.Title())
	m.SetBody(req.Body())
	m.SetCreated(req.Created())
	m.SetModified(req.Modified())
}

func (m *commentStruct) ToPrimary(ctx context.Context, req interface{}) (int64, error) {
	var resp int64
	return resp, nil
}
