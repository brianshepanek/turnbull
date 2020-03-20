package interactor

import (
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
)

type fooInteractor struct {
	fooInteractorStruct
}
type FooInteractor interface {
	fooInteractorInterface
}

func NewFooInteractor(r repository.FooRepository, p presenter.FooPresenter) FooInteractor {
	return &fooInteractor{fooInteractorStruct{r, p}}
}
