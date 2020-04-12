package entity

type enhancedAccount struct {
	*enhancedAccountStruct
}

type enhancedAccounts struct {
	enhancedAccountsStruct
}

type EnhancedAccount interface {
	enhancedAccountInterface
}
type EnhancedAccounts interface {
	enhancedAccountsInterface
}

func NewEnhancedAccount() EnhancedAccount {
	return newEnhancedAccount()
}

func NewEnhancedAccounts() EnhancedAccounts {
	return &enhancedAccounts{}
}
