package entity

type comment struct {
	commentStruct
}

type comments struct {
	commentsStruct
}

type Comment interface {
	commentInterface
}
type Comments interface {
	commentsInterface
}

func NewComment() Comment {
	return &comment{}
}

func NewComments() Comments {
	return &comments{}
}
