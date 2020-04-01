package controller

import interactor "github.com/brianshepanek/turnbull/turnbull/output/usecase/interactor"

type httpPostController struct {
	httpPostControllerStruct
}
type HttpPostController interface {
	httpPostControllerInterface
}

func New(interactor interactor.PostInteractor) HttpPostController {
	return &httpPostController{httpPostControllerStruct{interactor}}
}
