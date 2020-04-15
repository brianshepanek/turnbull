package registry

import (
	default1 "github.com/brianshepanek/turnbull/turnbull/output/interface/presenter/message/default"
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
)

type messageDefaultPresenterRegistry struct{}

func (r *registry) RegisterDefaultMessagePresenter() {}
func (r *registry) newDefaultMessagePresenter() presenter.MessagePresenter {
	return default1.New()
}
