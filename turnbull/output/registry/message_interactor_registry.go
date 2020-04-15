package registry

import interactor "github.com/brianshepanek/turnbull/turnbull/output/usecase/interactor"

func (r *registry) newMessageInteractor() interactor.MessageInteractor {
	return interactor.NewMessageInteractor(r.newMongoMessageRepository(), r.newDefaultMessagePresenter(), r.newAccountInteractor(), r.newChannelInteractor(), r.newUserInteractor())
}
