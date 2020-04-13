package registry

import (
	default1 "github.com/brianshepanek/turnbull/turnbull/output/interface/presenter/user/default"
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
)

type userDefaultPresenterRegistry struct{}

func (r *registry) RegisterDefaultUserPresenter() {}
func (r *registry) newDefaultUserPresenter() presenter.UserPresenter {
	return default1.New()
}
