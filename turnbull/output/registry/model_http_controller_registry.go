package registry

import (
	http "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/model/http"
	interactor "github.com/brianshepanek/turnbull/turnbull/output/usecase/interactor"
)

type modelHttpControllerRegistry struct{}

func (r *registry) RegisterHttpModelController() {}
func (r *registry) NewHttpModelController() http.HttpModelController {
	return http.New(interactor.NewModelInteractor(r.newMongoModelRepository(), r.newDefaultModelPresenter()))
}
