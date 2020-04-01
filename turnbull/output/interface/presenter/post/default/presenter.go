package presenter

type defaultPostPresenter struct {
	defaultPostPresenterStruct
}

func New() *defaultPostPresenterStruct {
	return &defaultPostPresenterStruct{}
}
