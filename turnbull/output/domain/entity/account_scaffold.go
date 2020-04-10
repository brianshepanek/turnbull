package entity

type accountStruct struct {
	*model
	*lucky
	name  *string
	email *string
}

func newAccount() *account {
	return &account{accountStruct: newAccountStruct()}
}

func newAccountStruct() *accountStruct {
	return &accountStruct{
		lucky: newLucky(),
		model: newModel(),
	}
}

type accountsStruct []accountInterface

type accountInterface interface {
	Name() *string
	Email() *string
	SetName(name *string)
	SetEmail(email *string)
	SetAll(req accountInterface)
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

func (m *accountStruct) Email() *string {
	return m.email
}

func (m *accountStruct) SetName(name *string) {
	m.name = name
}

func (m *accountStruct) SetEmail(email *string) {
	m.email = email
}

func (m *accountStruct) SetAll(req accountInterface) {
	m.SetName(req.Name())
	m.SetEmail(req.Email())
}
