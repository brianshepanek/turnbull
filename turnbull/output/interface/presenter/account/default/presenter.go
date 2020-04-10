package presenter

type defaultAccountPresenter struct {
	defaultAccountPresenterStruct
}

func New() *defaultAccountPresenter {
	return &defaultAccountPresenter{}
}
