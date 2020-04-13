package entity

type user struct {
	*userStruct
}

type users struct {
	usersStruct
}

type User interface {
	userInterface
}
type Users interface {
	usersInterface
}

func NewUser() User {
	return newUser()
}

func NewUsers() Users {
	return &users{}
}
