package entity

import "context"

type userStruct struct {
	*model
	accountId *int64
	firstName *string
	lastName  *string
	email     *string
}

func newUser() *user {
	return &user{userStruct: newUserStruct()}
}

func newUserStruct() *userStruct {
	return &userStruct{model: newModel()}
}

type usersStruct []userInterface

type userInterface interface {
	Model
	AccountId() *int64
	FirstName() *string
	LastName() *string
	Email() *string
	SetAccountId(accountId *int64)
	SetFirstName(firstName *string)
	SetLastName(lastName *string)
	SetEmail(email *string)
	BeforeRead(ctx context.Context) error
	BeforeAdd(ctx context.Context) error
}
type usersInterface interface {
	Len() int
	Append(req userInterface)
	Elements() []userInterface
}

func (m *usersStruct) Len() int {
	if m != nil {
		return len(*m)
	}
	return 0
}
func (m *usersStruct) Append(req userInterface) {
	if m != nil {
		*m = append(*m, req)
	}
}
func (m *usersStruct) Elements() []userInterface {
	return *m
}
func (m *userStruct) AccountId() *int64 {
	return m.accountId
}

func (m *userStruct) FirstName() *string {
	return m.firstName
}

func (m *userStruct) LastName() *string {
	return m.lastName
}

func (m *userStruct) Email() *string {
	return m.email
}

func (m *userStruct) SetAccountId(accountId *int64) {
	m.accountId = accountId
}

func (m *userStruct) SetFirstName(firstName *string) {
	m.firstName = firstName
}

func (m *userStruct) SetLastName(lastName *string) {
	m.lastName = lastName
}

func (m *userStruct) SetEmail(email *string) {
	m.email = email
}

func (m *userStruct) BeforeRead(ctx context.Context) error {
	return nil
}

func (m *userStruct) BeforeAdd(ctx context.Context) error {
	return nil
}
