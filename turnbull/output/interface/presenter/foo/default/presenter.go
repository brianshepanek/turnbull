package presenter

type defaultFooPresenter struct {
	defaultFooPresenterStruct
}

func New() *defaultFooPresenterStruct {
	return &defaultFooPresenterStruct{}
}
