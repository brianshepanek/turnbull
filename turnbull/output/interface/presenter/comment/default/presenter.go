package presenter

type defaultCommentPresenter struct {
	defaultCommentPresenterStruct
}

func New() *defaultCommentPresenterStruct {
	return &defaultCommentPresenterStruct{}
}
