package controller

import interactor "github.com/brianshepanek/turnbull/turnbull/output/usecase/interactor"

type httpAccountController struct {
	httpAccountControllerStruct
}
type HttpAccountController interface {
	httpAccountControllerInterface
}

func New(interactor interactor.AccountInteractor) HttpAccountController {
	return &httpAccountController{httpAccountControllerStruct{interactor}}
}
