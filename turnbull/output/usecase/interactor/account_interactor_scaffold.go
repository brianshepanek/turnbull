package interactor

import (
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
)

type accountInteractorStruct struct {
	repository repository.AccountRepository
	presenter  presenter.AccountPresenter
}
type accountInteractorInterface interface{}
