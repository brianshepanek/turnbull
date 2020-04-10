package controller

import interactor "github.com/brianshepanek/turnbull/turnbull/output/usecase/interactor"

type httpAccountControllerStruct struct {
	interactor interactor.AccountInteractor
}
type httpAccountControllerInterface interface{}
