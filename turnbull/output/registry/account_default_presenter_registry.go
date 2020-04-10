package registry

import (
	default1 "github.com/brianshepanek/turnbull/turnbull/output/interface/presenter/account/default"
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
)

type accountDefaultPresenterRegistry struct{}

func (r *registry) RegisterDefaultAccountPresenter() {}
func (r *registry) newDefaultAccountPresenter() presenter.AccountPresenter {
	return default1.New()
}
