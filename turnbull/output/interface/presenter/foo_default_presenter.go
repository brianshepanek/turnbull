package presenter

type defaultFooPresenter struct {
	defaultFooPresenterStruct
}

func NewDefaultFooPresenter() *defaultFooPresenterStruct {
	return &defaultFooPresenterStruct{}
}
