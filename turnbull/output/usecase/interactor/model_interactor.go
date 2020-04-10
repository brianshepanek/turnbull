package interactor

import (
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
)

type modelInteractor struct {
	modelInteractorStruct
}
type ModelInteractor interface {
	modelInteractorInterface
}

func NewModelInteractor(r repository.ModelRepository, p presenter.ModelPresenter) ModelInteractor {
	return &modelInteractor{modelInteractorStruct{r, p}}
}
