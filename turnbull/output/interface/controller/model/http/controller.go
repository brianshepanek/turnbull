package controller

import interactor "github.com/brianshepanek/turnbull/turnbull/output/usecase/interactor"

type httpModelController struct {
	httpModelControllerStruct
}
type HttpModelController interface {
	httpModelControllerInterface
}

func New(interactor interactor.ModelInteractor) HttpModelController {
	return &httpModelController{httpModelControllerStruct{interactor}}
}
