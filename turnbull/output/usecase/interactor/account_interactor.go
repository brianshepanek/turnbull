package interactor

import (
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
)

type accountInteractor struct {
	accountInteractorStruct
}
type AccountInteractor interface {
	accountInteractorInterface
}

func NewAccountInteractor(r repository.AccountRepository, p presenter.AccountPresenter) AccountInteractor {
	return &accountInteractor{accountInteractorStruct{r, p}}
}
