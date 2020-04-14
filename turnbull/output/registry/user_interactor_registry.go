package registry

import interactor "github.com/brianshepanek/turnbull/turnbull/output/usecase/interactor"

func (r *registry) newUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.newMongoUserRepository(), r.newDefaultUserPresenter(), r.newAccountInteractor())
}
