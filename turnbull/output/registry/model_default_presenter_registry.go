package registry

import (
	default1 "github.com/brianshepanek/turnbull/turnbull/output/interface/presenter/model/default"
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
)

type modelDefaultPresenterRegistry struct{}

func (r *registry) RegisterDefaultModelPresenter() {}
func (r *registry) newDefaultModelPresenter() presenter.ModelPresenter {
	return default1.New()
}
