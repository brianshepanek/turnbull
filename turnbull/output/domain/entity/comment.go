package entity

type Comment struct {
	commentStruct
}

func NewComment() *Comment {
	return &Comment{}
}

func NewComments() *[]Comment {
	return &[]Comment{}
}
