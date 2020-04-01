package controller

import interactor "github.com/brianshepanek/turnbull/turnbull/output/usecase/interactor"

type httpCommentController struct {
	httpCommentControllerStruct
}
type HttpCommentController interface {
	httpCommentControllerInterface
}

func New(interactor interactor.CommentInteractor) HttpCommentController {
	return &httpCommentController{httpCommentControllerStruct{interactor}}
}
