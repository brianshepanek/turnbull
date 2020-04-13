package registry

import interactor "github.com/brianshepanek/turnbull/turnbull/output/usecase/interactor"

func (r *registry) newAccountInteractor() interactor.AccountInteractor {
	return interactor.NewAccountInteractor(r.newMongoAccountRepository(), r.newDefaultAccountPresenter())
}
