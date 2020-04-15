package interactor

import (
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
)

type channelInteractor struct {
	channelInteractorStruct
}
type ChannelInteractor interface {
	channelInteractorInterface
}

func NewChannelInteractor(repository repository.ChannelRepository, presenter presenter.ChannelPresenter, accountInteractor AccountInteractor) ChannelInteractor {
	return &channelInteractor{channelInteractorStruct{repository, presenter, accountInteractor}}
}
