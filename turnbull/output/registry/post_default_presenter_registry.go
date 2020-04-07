package registry

import (
	default1 "github.com/brianshepanek/turnbull/turnbull/output/interface/presenter/post/default"
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
)

type postDefaultPresenterRegistry struct{}

func (r *registry) RegisterDefaultPostPresenter() {}
func (r *registry) newDefaultPostPresenter() presenter.PostPresenter {
	return default1.New()
}
