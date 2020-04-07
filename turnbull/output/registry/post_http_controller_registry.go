package registry

import (
	http "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/post/http"
	interactor "github.com/brianshepanek/turnbull/turnbull/output/usecase/interactor"
)

type postHttpControllerRegistry struct{}

func (r *registry) RegisterPostHttpControllerRegistry() {}
func (r *registry) NewPostHttpController() http.HttpPostController {
	return http.New(interactor.NewPostInteractor(r.newPostMongoRepository(), r.newPostDefaultPresenter()))
}
