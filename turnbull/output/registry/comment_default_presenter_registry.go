package registry

import (
	default1 "github.com/brianshepanek/turnbull/turnbull/output/interface/presenter/comment/default"
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
)

type commentDefaultPresenterRegistry struct{}

func (r *registry) RegisterDefaultCommentPresenter() {}
func (r *registry) newDefaultCommentPresenter() presenter.CommentPresenter {
	return default1.New()
}
