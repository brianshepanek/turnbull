package controller

import interactor "github.com/brianshepanek/turnbull/turnbull/output/usecase/interactor"

type httpUserController struct {
	httpUserControllerStruct
}
type HttpUserController interface {
	httpUserControllerInterface
}

func New(interactor interactor.UserInteractor) HttpUserController {
	return &httpUserController{httpUserControllerStruct{interactor}}
}
