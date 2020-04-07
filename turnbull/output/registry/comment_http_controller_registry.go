package registry

import (
	http "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/comment/http"
	interactor "github.com/brianshepanek/turnbull/turnbull/output/usecase/interactor"
)

type commentHttpControllerRegistry struct{}

func (r *registry) RegisterCommentHttpControllerRegistry() {}
func (r *registry) NewCommentHttpController() http.HttpCommentController {
	return http.New(interactor.NewCommentInteractor(r.newCommentMysqlRepository(), r.newCommentDefaultPresenter()))
}
