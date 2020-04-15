package controller

import interactor "github.com/brianshepanek/turnbull/turnbull/output/usecase/interactor"

type httpChannelController struct {
	httpChannelControllerStruct
}
type HttpChannelController interface {
	httpChannelControllerInterface
}

func New(interactor interactor.ChannelInteractor) HttpChannelController {
	return &httpChannelController{httpChannelControllerStruct{interactor}}
}
