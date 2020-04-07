package registry

import (
	default1 "github.com/brianshepanek/turnbull/turnbull/output/interface/presenter/comment/default"
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
)

type commentDefaultPresenterRegistry struct{}

func (r *registry) RegisterCommentDefaultPresenterRegistry() {}
func (r *registry) newCommentDefaultPresenterRegistry() presenter.CommentPresenter {
	return default1.New()
}
