package interactor

import (
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
)

type luckyInteractorStruct struct {
	repository repository.LuckyRepository
	presenter  presenter.LuckyPresenter
}
type luckyInteractorInterface interface{}
