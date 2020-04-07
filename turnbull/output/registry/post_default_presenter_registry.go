package registry

import (
	default1 "github.com/brianshepanek/turnbull/turnbull/output/interface/presenter/post/default"
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
)

type postDefaultPresenterRegistry struct{}

func (r *registry) RegisterPostDefaultPresenterRegistry() {}
func (r *registry) newPostDefaultPresenterRegistry() presenter.PostPresenter {
	return default1.New()
}
