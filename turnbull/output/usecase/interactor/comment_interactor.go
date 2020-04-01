package interactor

import (
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
)

type commentInteractor struct {
	commentInteractorStruct
}
type CommentInteractor interface {
	commentInteractorInterface
}

func NewCommentInteractor(r repository.CommentRepository, p presenter.CommentPresenter) CommentInteractor {
	return &commentInteractor{commentInteractorStruct{r, p}}
}
