package entity

import (
	"context"
	"time"
)

type modelStruct struct {
	id       *int64
	created  *time.Time
	modified *time.Time
}

func newModel() *model {
	return &model{modelStruct: newModelStruct()}
}

func newModelStruct() *modelStruct {
	return &modelStruct{}
}

type modelsStruct []modelInterface

type modelInterface interface {
	Id() *int64
	Created() *time.Time
	Modified() *time.Time
	SetId(id *int64)
	SetCreated(created *time.Time)
	SetModified(modified *time.Time)
	SetAll(req modelInterface)
	ToPrimary(ctx context.Context, req interface{}) (int64, error)
}
type modelsInterface interface {
	Len() int
	Append(req modelInterface)
	Elements() []modelInterface
}

func (m *modelsStruct) Len() int {
	if m != nil {
		return len(*m)
	}
	return 0
}
func (m *modelsStruct) Append(req modelInterface) {
	if m != nil {
		*m = append(*m, req)
	}
}
func (m *modelsStruct) Elements() []modelInterface {
	return *m
}
func (m *modelStruct) Id() *int64 {
	return m.id
}

func (m *modelStruct) Created() *time.Time {
	return m.created
}

func (m *modelStruct) Modified() *time.Time {
	return m.modified
}

func (m *modelStruct) SetId(id *int64) {
	m.id = id
}

func (m *modelStruct) SetCreated(created *time.Time) {
	m.created = created
}

func (m *modelStruct) SetModified(modified *time.Time) {
	m.modified = modified
}

func (m *modelStruct) SetAll(req modelInterface) {
	m.SetId(req.Id())
	m.SetCreated(req.Created())
	m.SetModified(req.Modified())
}

func (m *modelStruct) ToPrimary(ctx context.Context, req interface{}) (int64, error) {
	var resp int64
	return resp, nil
}
