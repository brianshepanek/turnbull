package registry

import interactor "github.com/brianshepanek/turnbull/turnbull/output/usecase/interactor"

func (r *registry) newChannelInteractor() interactor.ChannelInteractor {
	return interactor.NewChannelInteractor(r.newMongoChannelRepository(), r.newDefaultChannelPresenter(), r.newAccountInteractor())
}
