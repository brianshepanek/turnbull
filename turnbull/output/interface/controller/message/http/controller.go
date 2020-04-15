package controller

import interactor "github.com/brianshepanek/turnbull/turnbull/output/usecase/interactor"

type httpMessageController struct {
	httpMessageControllerStruct
}
type HttpMessageController interface {
	httpMessageControllerInterface
}

func New(interactor interactor.MessageInteractor) HttpMessageController {
	return &httpMessageController{httpMessageControllerStruct{interactor}}
}
