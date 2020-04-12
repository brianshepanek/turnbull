package controller

import interactor "github.com/brianshepanek/turnbull/turnbull/output/usecase/interactor"

type httpEnhancedAccountController struct {
	httpEnhancedAccountControllerStruct
}
type HttpEnhancedAccountController interface {
	httpEnhancedAccountControllerInterface
}

func New(interactor interactor.EnhancedAccountInteractor) HttpEnhancedAccountController {
	return &httpEnhancedAccountController{httpEnhancedAccountControllerStruct{interactor}}
}
