package entity

import (
	"context"
	"time"
)

type accountStruct struct {
	id       *int64
	name     *string
	created  *time.Time
	modified *time.Time
}

func newAccount() *account {
	return &account{accountStruct: newAccountStruct()}
}

func newAccountStruct() *accountStruct {
	return &accountStruct{}
}

type accountsStruct []accountInterface

type accountInterface interface {
	Id() *int64
	Name() *string
	Created() *time.Time
	Modified() *time.Time
	SetId(id *int64)
	SetName(name *string)
	SetCreated(created *time.Time)
	SetModified(modified *time.Time)
	BeforeRead(ctx context.Context) error
	BeforeAdd(ctx context.Context) error
	SetAll(req accountInterface)
	ToPrimary(ctx context.Context, req interface{}) (int64, error)
}
type accountsInterface interface {
	Len() int
	Append(req accountInterface)
	Elements() []accountInterface
}

func (m *accountsStruct) Len() int {
	if m != nil {
		return len(*m)
	}
	return 0
}
func (m *accountsStruct) Append(req accountInterface) {
	if m != nil {
		*m = append(*m, req)
	}
}
func (m *accountsStruct) Elements() []accountInterface {
	return *m
}
func (m *accountStruct) Id() *int64 {
	return m.id
}

func (m *accountStruct) Name() *string {
	return m.name
}

func (m *accountStruct) Created() *time.Time {
	return m.created
}

func (m *accountStruct) Modified() *time.Time {
	return m.modified
}

func (m *accountStruct) SetId(id *int64) {
	m.id = id
}

func (m *accountStruct) SetName(name *string) {
	m.name = name
}

func (m *accountStruct) SetCreated(created *time.Time) {
	m.created = created
}

func (m *accountStruct) SetModified(modified *time.Time) {
	m.modified = modified
}

func (m *accountStruct) SetAll(req accountInterface) {
	m.SetId(req.Id())
	m.SetName(req.Name())
	m.SetCreated(req.Created())
	m.SetModified(req.Modified())
}

func (m *accountStruct) BeforeRead(ctx context.Context) error {
	return nil
}

func (m *accountStruct) BeforeAdd(ctx context.Context) error {
	return nil
}

func (m *accountStruct) ToPrimary(ctx context.Context, req interface{}) (int64, error) {
	var resp int64
	return resp, nil
}
