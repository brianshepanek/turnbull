package interactor

import (
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
)

type luckyInteractor struct {
	luckyInteractorStruct
}
type LuckyInteractor interface {
	luckyInteractorInterface
}

func NewLuckyInteractor(r repository.LuckyRepository, p presenter.LuckyPresenter) LuckyInteractor {
	return &luckyInteractor{luckyInteractorStruct{r, p}}
}
