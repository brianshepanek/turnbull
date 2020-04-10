package presenter

type defaultModelPresenter struct {
	defaultModelPresenterStruct
}

func New() *defaultModelPresenter {
	return &defaultModelPresenter{}
}
