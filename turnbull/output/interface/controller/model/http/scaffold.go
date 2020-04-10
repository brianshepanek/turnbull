package controller

import interactor "github.com/brianshepanek/turnbull/turnbull/output/usecase/interactor"

type httpModelControllerStruct struct {
	interactor interactor.ModelInteractor
}
type httpModelControllerInterface interface{}
