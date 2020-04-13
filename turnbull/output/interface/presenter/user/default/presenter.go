package presenter

type defaultUserPresenter struct {
	defaultUserPresenterStruct
}

func New() *defaultUserPresenter {
	return &defaultUserPresenter{}
}
