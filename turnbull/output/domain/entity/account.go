package entity

type account struct {
	*accountStruct
}

type accounts struct {
	accountsStruct
}

type Account interface {
	accountInterface
}
type Accounts interface {
	accountsInterface
}

func NewAccount() Account {
	return newAccount()
}

func NewAccounts() Accounts {
	return &accounts{}
}
