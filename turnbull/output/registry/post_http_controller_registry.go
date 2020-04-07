package registry

import (
	http "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/post/http"
	interactor "github.com/brianshepanek/turnbull/turnbull/output/usecase/interactor"
)

type postHttpControllerRegistry struct{}

func (r *registry) RegisterHttpPostController() {}
func (r *registry) NewHttpPostController() http.HttpPostController {
	return http.New(interactor.NewPostInteractor(r.newMongoPostRepository(), r.newDefaultPostPresenter()))
}
