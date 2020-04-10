package controller

import interactor "github.com/brianshepanek/turnbull/turnbull/output/usecase/interactor"

type httpLuckyControllerStruct struct {
	interactor interactor.LuckyInteractor
}
type httpLuckyControllerInterface interface{}
