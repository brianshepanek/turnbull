package interactor

import (
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
)

type userInteractor struct {
	userInteractorStruct
}
type UserInteractor interface {
	userInteractorInterface
}

func NewUserInteractor(repository repository.UserRepository, presenter presenter.UserPresenter, accountInteractor AccountInteractor) UserInteractor {
	return &userInteractor{userInteractorStruct{repository, presenter, accountInteractor}}
}
