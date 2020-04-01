package interactor

import (
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
)

type postInteractor struct {
	postInteractorStruct
}
type PostInteractor interface {
	postInteractorInterface
}

func NewPostInteractor(r repository.PostRepository, p presenter.PostPresenter) PostInteractor {
	return &postInteractor{postInteractorStruct{r, p}}
}
