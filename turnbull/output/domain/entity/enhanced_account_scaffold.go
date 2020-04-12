package entity

import "context"

type enhancedAccountStruct struct {
	*account
	enhancement *string
}

func newEnhancedAccount() *enhancedAccount {
	return &enhancedAccount{enhancedAccountStruct: newEnhancedAccountStruct()}
}

func newEnhancedAccountStruct() *enhancedAccountStruct {
	return &enhancedAccountStruct{account: newAccount()}
}

type enhancedAccountsStruct []enhancedAccountInterface

type enhancedAccountInterface interface {
	Account
	Enhancement() *string
	SetEnhancement(enhancement *string)
	BeforeRead(ctx context.Context) error
	BeforeAdd(ctx context.Context) error
}
type enhancedAccountsInterface interface {
	Len() int
	Append(req enhancedAccountInterface)
	Elements() []enhancedAccountInterface
}

func (m *enhancedAccountsStruct) Len() int {
	if m != nil {
		return len(*m)
	}
	return 0
}
func (m *enhancedAccountsStruct) Append(req enhancedAccountInterface) {
	if m != nil {
		*m = append(*m, req)
	}
}
func (m *enhancedAccountsStruct) Elements() []enhancedAccountInterface {
	return *m
}
func (m *enhancedAccountStruct) Enhancement() *string {
	return m.enhancement
}

func (m *enhancedAccountStruct) SetEnhancement(enhancement *string) {
	m.enhancement = enhancement
}

func (m *enhancedAccountStruct) BeforeRead(ctx context.Context) error {
	return nil
}

func (m *enhancedAccountStruct) BeforeAdd(ctx context.Context) error {
	return nil
}
