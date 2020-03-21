package controller

import interactor "github.com/brianshepanek/turnbull/turnbull/output/usecase/interactor"

type httpFooController struct {
	httpFooControllerStruct
}
type HttpFooController interface {
	httpFooControllerInterface
}

func NewHttpFooController(interactor interactor.FooInteractor) HttpFooController {
	return &httpFooController{httpFooControllerStruct{interactor}}
}
