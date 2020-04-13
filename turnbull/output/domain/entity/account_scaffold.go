package entity

import "context"

type accountStruct struct {
	*model
	name *string
}

func newAccount() *account {
	return &account{accountStruct: newAccountStruct()}
}

func newAccountStruct() *accountStruct {
	return &accountStruct{model: newModel()}
}

type accountsStruct []accountInterface

type accountInterface interface {
	Model
	Name() *string
	SetName(name *string)
	BeforeRead(ctx context.Context) error
	BeforeAdd(ctx context.Context) error
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
func (m *accountStruct) Name() *string {
	return m.name
}

func (m *accountStruct) SetName(name *string) {
	m.name = name
}

func (m *accountStruct) BeforeRead(ctx context.Context) error {
	return nil
}

func (m *accountStruct) BeforeAdd(ctx context.Context) error {
	return nil
}
