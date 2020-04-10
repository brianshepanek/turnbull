package controller

import interactor "github.com/brianshepanek/turnbull/turnbull/output/usecase/interactor"

type httpLuckyController struct {
	httpLuckyControllerStruct
}
type HttpLuckyController interface {
	httpLuckyControllerInterface
}

func New(interactor interactor.LuckyInteractor) HttpLuckyController {
	return &httpLuckyController{httpLuckyControllerStruct{interactor}}
}
