package presenter

type defaultMessagePresenter struct {
	defaultMessagePresenterStruct
}

func New() *defaultMessagePresenter {
	return &defaultMessagePresenter{}
}
