package interactor

import (
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
)

type messageInteractor struct {
	messageInteractorStruct
}
type MessageInteractor interface {
	messageInteractorInterface
}

func NewMessageInteractor(repository repository.MessageRepository, presenter presenter.MessagePresenter, accountInteractor AccountInteractor, channelInteractor ChannelInteractor, userInteractor UserInteractor) MessageInteractor {
	return &messageInteractor{messageInteractorStruct{repository, presenter, accountInteractor, channelInteractor, userInteractor}}
}
