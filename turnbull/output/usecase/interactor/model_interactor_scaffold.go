package interactor

import (
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
)

type modelInteractorStruct struct {
	repository repository.ModelRepository
	presenter  presenter.ModelPresenter
}
type modelInteractorInterface interface{}
