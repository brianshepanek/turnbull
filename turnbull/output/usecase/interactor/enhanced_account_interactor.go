package interactor

import (
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
)

type enhancedAccountInteractor struct {
	enhancedAccountInteractorStruct
}
type EnhancedAccountInteractor interface {
	enhancedAccountInteractorInterface
}

func NewEnhancedAccountInteractor(r repository.EnhancedAccountRepository, p presenter.EnhancedAccountPresenter) EnhancedAccountInteractor {
	return &enhancedAccountInteractor{enhancedAccountInteractorStruct{r, p}}
}
